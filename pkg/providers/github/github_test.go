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

package github

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRestClient(t *testing.T) {
	t.Parallel()

	client, err := NewRestClient(context.Background(), GitHubConfig{
		Token:    "token",
		Endpoint: "https://api.github.com",
	}, "")
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestNewGraphQLClient(t *testing.T) {
	t.Parallel()

	client, err := NewGraphQLClient(context.Background(), GitHubConfig{
		Token:    "token",
		Endpoint: "https://api.github.com/graphql",
	})

	assert.NoError(t, err)
	assert.NotNil(t, client)
}
