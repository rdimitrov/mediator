// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/stacklok/mediator/internal/db"
	evalerrors "github.com/stacklok/mediator/internal/engine/errors"
)

// createOrUpdateEvalStatusParams is a helper struct to pass parameters to createOrUpdateEvalStatus
// to avoid confusion with the parameters order. Since at the moment all our entities are bound to
// a repo and most policies are expecting a repo, the repoID parameter is mandatory. For entities
// other than artifacts, the artifactID should be 0 which is translated to NULL in the database.
type createOrUpdateEvalStatusParams struct {
	policyID       uuid.UUID
	repoID         uuid.UUID
	artifactID     *uuid.UUID
	ruleTypeEntity db.Entities
	ruleTypeID     uuid.UUID
	evalErr        error
}

func (e *Executor) createOrUpdateEvalStatus(
	ctx context.Context,
	params *createOrUpdateEvalStatusParams,
) error {
	if params == nil {
		return fmt.Errorf("createOrUpdateEvalStatusParams cannot be nil")
	}

	if errors.Is(params.evalErr, evalerrors.ErrEvaluationSkipSilently) {
		log.Printf("silent skip of rule %s for policy %s for entity %s in repo %s",
			params.ruleTypeID, params.policyID, params.ruleTypeEntity, params.repoID)
		return nil
	}

	var sqlArtifactID uuid.NullUUID
	if params.artifactID != nil {
		sqlArtifactID = uuid.NullUUID{
			UUID:  *params.artifactID,
			Valid: true,
		}
	}

	return e.querier.UpsertRuleEvaluationStatus(ctx, db.UpsertRuleEvaluationStatusParams{
		PolicyID: params.policyID,
		RepositoryID: uuid.NullUUID{
			UUID:  params.repoID,
			Valid: true,
		},
		ArtifactID: sqlArtifactID,
		Entity:     params.ruleTypeEntity,
		RuleTypeID: params.ruleTypeID,
		EvalStatus: errorAsEvalStatus(params.evalErr),
		Details:    errorAsDetails(params.evalErr),
	})
}

func errorAsEvalStatus(err error) db.EvalStatusTypes {
	if errors.Is(err, evalerrors.ErrEvaluationFailed) {
		return db.EvalStatusTypesFailure
	} else if errors.Is(err, evalerrors.ErrEvaluationSkipped) {
		return db.EvalStatusTypesSkipped
	} else if err != nil {
		return db.EvalStatusTypesError
	}
	return db.EvalStatusTypesSuccess
}

func errorAsDetails(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
