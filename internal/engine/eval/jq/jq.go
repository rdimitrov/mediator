// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.role/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package rule provides the CLI subcommand for managing rules

// Package jq provides the jq policy evaluator
package jq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	evalerrors "github.com/stacklok/mediator/internal/engine/errors"
	engif "github.com/stacklok/mediator/internal/engine/interfaces"
	"github.com/stacklok/mediator/internal/util"
	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
)

// Evaluator is an Evaluator that uses the jq library to evaluate rules
type Evaluator struct {
	assertions []*pb.RuleType_Definition_Eval_JQComparison
}

// NewJQEvaluator creates a new JQ rule data evaluator
func NewJQEvaluator(assertions []*pb.RuleType_Definition_Eval_JQComparison) (*Evaluator, error) {
	if len(assertions) == 0 {
		return nil, fmt.Errorf("missing jq assertions")
	}

	for idx := range assertions {
		a := assertions[idx]
		if a.Policy == nil {
			return nil, fmt.Errorf("missing policy accessor")
		}

		if a.Policy.Def == "" {
			return nil, fmt.Errorf("missing policy accessor definition")
		}

		if a.Ingested == nil {
			return nil, fmt.Errorf("missing data accessor")
		}

		if a.Ingested.Def == "" {
			return nil, fmt.Errorf("missing data accessor definition")
		}
	}

	return &Evaluator{
		assertions: assertions,
	}, nil
}

// Eval calls the jq library to evaluate the rule
func (jqe *Evaluator) Eval(ctx context.Context, pol map[string]any, res *engif.Result) error {
	if res.Object == nil {
		return fmt.Errorf("missing object")
	}
	obj := res.Object

	for idx := range jqe.assertions {
		var policyVal, dataVal any

		a := jqe.assertions[idx]
		policyVal, err := util.JQReadFrom[any](ctx, a.Policy.Def, pol)
		// we ignore util.ErrNoValueFound because we want to allow the JQ accessor to return the default value
		// which is fine for DeepEqual
		if err != nil && !errors.Is(err, util.ErrNoValueFound) {
			return fmt.Errorf("cannot get values from policy accessor: %w", err)
		}

		dataVal, err = util.JQReadFrom[any](ctx, a.Ingested.Def, obj)
		if err != nil && !errors.Is(err, util.ErrNoValueFound) {
			return fmt.Errorf("cannot get values from data accessor: %w", err)
		}

		// Deep compare
		if !reflect.DeepEqual(policyVal, dataVal) {
			msg := fmt.Sprintf("data does not match policy: for assertion %d, got %v, want %v",
				idx, dataVal, policyVal)

			marshalledAssertion, err := json.MarshalIndent(a, "", "  ")
			if err == nil {
				msg = fmt.Sprintf("%s\nassertion: %s", msg, string(marshalledAssertion))
			}

			return evalerrors.NewErrEvaluationFailed(msg)
		}
	}

	return nil
}
