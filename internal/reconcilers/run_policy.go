// Copyright 2023 Stacklok, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reconcilers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stacklok/mediator/internal/db"
	"github.com/stacklok/mediator/internal/engine"
	"github.com/stacklok/mediator/internal/util"
	pb "github.com/stacklok/mediator/pkg/api/protobuf/go/mediator/v1"
)

// PolicyInitEvent is an event that is sent to the reconciler topic
// when a new policy is created. It is used to initialize the policy
// by iterating over all registered entities for the relevant group
// and sending a policy evaluation event for each one.
type PolicyInitEvent struct {
	// Group is the group that the event is relevant to
	Group int32 `json:"group" validate:"gte=0"`
}

// NewPolicyInitMessage creates a new repos init event
func NewPolicyInitMessage(provider string, groupID int32) (*message.Message, error) {
	evt := &PolicyInitEvent{
		Group: groupID,
	}

	evtStr, err := json.Marshal(evt)
	if err != nil {
		return nil, fmt.Errorf("error marshalling init event: %w", err)
	}

	msg := message.NewMessage(uuid.New().String(), evtStr)
	msg.Metadata.Set("provider", provider)
	return msg, nil
}

// handlePolicyInitEvent handles a policy init event.
// It is responsible for iterating over all registered repositories
// for the group and sending a policy evaluation event for each one.
func (e *Reconciler) handlePolicyInitEvent(msg *message.Message) error {
	prov := msg.Metadata.Get("provider")

	var evt PolicyInitEvent
	if err := json.Unmarshal(msg.Payload, &evt); err != nil {
		return fmt.Errorf("error unmarshalling payload: %w", err)
	}

	// validate event
	validate := validator.New()
	if err := validate.Struct(evt); err != nil {
		// We don't return the event since there's no use
		// retrying it if it's invalid.
		log.Printf("error validating event: %v", err)
		return nil
	}

	provInfo, err := e.store.GetProviderByName(context.Background(), db.GetProviderByNameParams{
		Name:    prov,
		GroupID: evt.Group,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// We don't return the event since there's no use
			// retrying it if the provider doesn't exist.
			log.Printf("provider %s not found", prov)
			return nil
		}

		return fmt.Errorf("error getting provider: %w", err)
	}

	ectx := &engine.EntityContext{
		Group: engine.Group{
			ID: evt.Group,
		},
		Provider: engine.Provider{
			Name: provInfo.Name,
			ID:   provInfo.ID,
		},
	}

	ctx := msg.Context()
	log.Printf("handling policy init event for group %d", evt.Group)
	if err := e.publishPolicyInitEvents(ctx, ectx); err != nil {
		// We don't return an error since watermill will retry
		// the message.
		log.Printf("publishPolicyInitEvents: error publishing policy events: %v", err)
		return nil
	}

	return nil
}

func (s *Reconciler) publishPolicyInitEvents(
	ctx context.Context,
	ectx *engine.EntityContext,
) error {
	dbrepos, err := s.store.ListRegisteredRepositoriesByGroupIDAndProvider(ctx,
		db.ListRegisteredRepositoriesByGroupIDAndProviderParams{
			Provider: ectx.Provider.Name,
			GroupID:  ectx.Group.ID,
		})
	if err != nil {
		return fmt.Errorf("publishPolicyInitEvents: error getting registered repos: %v", err)
	}

	for _, dbrepo := range dbrepos {
		// protobufs are our API, so we always execute on these instead of the DB directly.
		repo := &pb.RepositoryResult{
			Owner:      dbrepo.RepoOwner,
			Repository: dbrepo.RepoName,
			RepoId:     dbrepo.RepoID,
			HookUrl:    dbrepo.WebhookUrl,
			DeployUrl:  dbrepo.DeployUrl,
			CloneUrl:   dbrepo.CloneUrl,
			CreatedAt:  timestamppb.New(dbrepo.CreatedAt),
			UpdatedAt:  timestamppb.New(dbrepo.UpdatedAt),
		}

		err := engine.NewEntityInfoWrapper().
			WithProvider(ectx.Provider.Name).
			WithGroupID(ectx.Group.ID).
			WithRepository(repo).
			WithRepositoryID(dbrepo.ID).
			Publish(s.evt)

		// This is a non-fatal error, so we'll just log it
		// and continue
		if err != nil {
			return fmt.Errorf("error publishing init event for repo %s: %v", dbrepo.ID, err)
		}
	}

	// after we've initialized repository policies, let's initialize artifacts
	// TODO(jakub): this should be done in an iterator of sorts
	for i := range dbrepos {
		pdb := &dbrepos[i]
		err := s.publishArtifactPolicyInitEvents(ctx, ectx, pdb)
		if err != nil {
			return fmt.Errorf("publishPolicyInitEvents: error publishing artifact events: %v", err)
		}
	}

	return nil
}

func (s *Reconciler) publishArtifactPolicyInitEvents(
	ctx context.Context,
	ectx *engine.EntityContext,
	dbrepo *db.Repository,
) error {
	dbArtifacts, err := s.store.ListArtifactsByRepoID(ctx, dbrepo.ID)
	if err != nil {
		return fmt.Errorf("error getting artifacts: %w", err)
	}
	if len(dbArtifacts) == 0 {
		zerolog.Ctx(ctx).Debug().Str("repository", dbrepo.ID.String()).Msgf("no artifacts found, skipping")
		return nil
	}
	for _, dbA := range dbArtifacts {
		// Get the artifact with all its versions as a protobuf
		pbArtifact, err := util.GetArtifactWithVersions(ctx, s.store, dbrepo.ID, dbA.ID)
		if err != nil {
			return fmt.Errorf("error getting artifact versions: %w", err)
		}

		err = engine.NewEntityInfoWrapper().
			WithProvider(ectx.Provider.Name).
			WithGroupID(ectx.Group.ID).
			WithArtifact(pbArtifact).
			WithRepositoryID(dbrepo.ID).
			WithArtifactID(dbA.ID).
			Publish(s.evt)

		// This is a non-fatal error, so we'll just log it
		// and continue
		if err != nil {
			log.Printf("error publishing init event for repo %s: %v", dbrepo.ID, err)
			continue
		}
	}
	return nil
}
