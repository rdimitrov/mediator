//
// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// NOTE: This file is for stubbing out client code for proof of concept
// purposes. It will / should be removed in the future.
// Until then, it is not covered by unit tests and should not be used
// It does make a good example of how to use the generated client code
// for others to use as a reference.

package auth

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/stacklok/mediator/internal/util"
	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
)

// Auth_revokeproviderCmd represents the auth revoke command
var Auth_revokeproviderCmd = &cobra.Command{
	Use:   "revoke_provider",
	Short: "Revoke access tokens for provider",
	Long:  `It can revoke access tokens for specific provider.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			fmt.Fprintf(os.Stderr, "Error binding flags: %s\n", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// check if we need to revoke all tokens or the user one
		all := util.GetConfigValue("all", "all", cmd, false).(bool)
		group := viper.GetInt32("group-id")
		provider := util.GetConfigValue("provider", "provider", cmd, "").(string)

		if all && group != 0 {
			fmt.Fprintf(os.Stderr, "Error: you can't use --all and --group-id together\n")
			os.Exit(1)
		}

		conn, err := util.GrpcForCommand(cmd)
		util.ExitNicelyOnError(err, "Error getting grpc connection")
		defer conn.Close()

		ctx, cancel := util.GetAppContext()
		defer cancel()
		client := pb.NewOAuthServiceClient(conn)
		if all {
			result, err := client.RevokeOauthTokens(ctx, &pb.RevokeOauthTokensRequest{Provider: provider})
			util.ExitNicelyOnError(err, "Error revoking tokens")
			cmd.Println("Revoked a total of ", result.RevokedTokens, " tokens")
		} else {
			_, err := client.RevokeOauthGroupToken(ctx, &pb.RevokeOauthGroupTokenRequest{Provider: provider, GroupId: group})
			util.ExitNicelyOnError(err, "Error revoking tokens")
			if group == 0 {
				cmd.Println("Revoked token for default group")
			} else {
				cmd.Println("Revoked token for group ", group)
			}
		}
	},
}

func init() {
	AuthCmd.AddCommand(Auth_revokeproviderCmd)
	Auth_revokeproviderCmd.Flags().StringP("provider", "n", "", "Name for the provider to revoke tokens for")
	Auth_revokeproviderCmd.Flags().BoolP("all", "a", false, "Revoke all tokens")
}
