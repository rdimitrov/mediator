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

package rest

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
	ghclient "github.com/stacklok/mediator/pkg/providers/github"
)

func TestNewRestRuleDataIngest(t *testing.T) {
	t.Parallel()

	type args struct {
		restCfg *pb.RestType
		cli     ghclient.RestAPI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid rest",
			args: args{
				restCfg: &pb.RestType{
					Endpoint: "https://api.github.com/repos/Foo/Bar",
				},
				cli: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid template",
			args: args{
				restCfg: &pb.RestType{
					Endpoint: "{{",
				},
			},
			wantErr: true,
		},
		{
			name: "empty endpoint",
			args: args{
				restCfg: &pb.RestType{
					Endpoint: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewRestRuleDataIngest(tt.args.restCfg, tt.args.cli)
			if tt.wantErr {
				require.Error(t, err, "expected error")
				require.Nil(t, got, "expected nil")
				return
			}

			require.NoError(t, err, "unexpected error")
			require.NotNil(t, got, "expected non-nil")
		})
	}
}
