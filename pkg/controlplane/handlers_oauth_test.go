// Copyright 2023 Stacklok, Inc
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

package controlplane

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc/codes"

	mockdb "github.com/stacklok/mediator/database/mock"
	"github.com/stacklok/mediator/internal/config"
	"github.com/stacklok/mediator/pkg/auth"
	"github.com/stacklok/mediator/pkg/db"
	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
	ghclient "github.com/stacklok/mediator/pkg/providers/github"
)

func TestNewOAuthConfig(t *testing.T) {
	t.Parallel()

	// Test with CLI set
	cfg, err := auth.NewOAuthConfig("google", true)
	if err != nil {
		t.Errorf("Error in newOAuthConfig: %v", err)
	}

	if cfg.Endpoint != google.Endpoint {
		t.Errorf("Unexpected endpoint: %v", cfg.Endpoint)
	}

	// Test with CLI set
	cfg, err = auth.NewOAuthConfig("github", true)
	if err != nil {
		t.Errorf("Error in newOAuthConfig: %v", err)
	}

	if cfg.Endpoint != github.Endpoint {
		t.Errorf("Unexpected endpoint: %v", cfg.Endpoint)
	}

	// Test with CLI set
	cfg, err = auth.NewOAuthConfig("google", false)
	if err != nil {
		t.Errorf("Error in newOAuthConfig: %v", err)
	}

	if cfg.Endpoint != google.Endpoint {
		t.Errorf("Unexpected endpoint: %v", cfg.Endpoint)
	}

	// Test with CLI set
	cfg, err = auth.NewOAuthConfig("github", false)
	if err != nil {
		t.Errorf("Error in newOAuthConfig: %v", err)
	}

	if cfg.Endpoint != github.Endpoint {
		t.Errorf("Unexpected endpoint: %v", cfg.Endpoint)
	}

	_, err = auth.NewOAuthConfig("invalid", true)
	if err == nil {
		t.Errorf("Expected error in newOAuthConfig, but got nil")
	}
}

func TestGetAuthorizationURL(t *testing.T) {
	t.Parallel()

	state := "test"
	grpID := sql.NullInt32{Int32: 1, Valid: true}
	port := sql.NullInt32{Int32: 8080, Valid: true}

	testCases := []struct {
		name               string
		req                *pb.GetAuthorizationURLRequest
		buildStubs         func(store *mockdb.MockStore)
		checkResponse      func(t *testing.T, res *pb.GetAuthorizationURLResponse, err error)
		expectedStatusCode codes.Code
	}{
		{
			name: "Success",
			req: &pb.GetAuthorizationURLRequest{
				Provider: "github",
				Port:     8080,
				Cli:      true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSessionState(gomock.Any(), gomock.Any()).
					Return(db.SessionStore{
						GrpID:        grpID,
						Port:         port,
						SessionState: state,
					}, nil)
				store.EXPECT().
					DeleteSessionStateByGroupID(gomock.Any(), gomock.Any()).
					Return(nil)
			},

			checkResponse: func(t *testing.T, res *pb.GetAuthorizationURLResponse, err error) {
				t.Helper()

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if res.Url == "" {
					t.Errorf("Unexpected response from GetAuthorizationURL: %v", res)
				}
			},

			expectedStatusCode: codes.OK,
		},
	}

	// Create a new context and set the claims value
	ctx := context.WithValue(context.Background(), auth.TokenInfoKey, auth.UserClaims{
		UserId:         1,
		OrganizationId: 1,
		GroupIds:       []int32{1},
		Roles: []auth.RoleInfo{
			{RoleID: 1, IsAdmin: true, GroupID: 0, OrganizationID: 1}},
	})

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := Server{store: store}

			res, err := server.GetAuthorizationURL(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestRevokeOauthTokens_gRPC(t *testing.T) {
	t.Parallel()

	ctx := context.WithValue(context.Background(), auth.TokenInfoKey, auth.UserClaims{
		UserId:         1,
		OrganizationId: 1,
		GroupIds:       []int32{1},
		Roles: []auth.RoleInfo{
			{RoleID: 1, IsAdmin: true, GroupID: 0, OrganizationID: 1}},
	})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdb.NewMockStore(ctrl)
	mockStore.EXPECT().GetAccessTokenByProvider(gomock.Any(), gomock.Any())

	server, err := NewServer(mockStore, &config.Config{})
	require.NoError(t, err, "failed to create test server")

	res, err := server.RevokeOauthTokens(ctx, &pb.RevokeOauthTokensRequest{Provider: ghclient.Github})

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func RevokeOauthGroupToken_gRPC(t *testing.T) {
	t.Helper()

	ctx := context.WithValue(context.Background(), auth.TokenInfoKey, auth.UserClaims{
		UserId:         1,
		OrganizationId: 1,
		GroupIds:       []int32{1},
		Roles: []auth.RoleInfo{
			{RoleID: 1, IsAdmin: true, GroupID: 0, OrganizationID: 1}},
	})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdb.NewMockStore(ctrl)
	mockStore.EXPECT().GetAccessTokenByGroupID(gomock.Any(), gomock.Any())

	server, err := NewServer(mockStore, &config.Config{})
	require.NoError(t, err, "failed to create test server")

	res, err := server.RevokeOauthGroupToken(ctx, &pb.RevokeOauthGroupTokenRequest{Provider: ghclient.Github, GroupId: 1})

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
