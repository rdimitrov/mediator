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

package rule_type

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/stacklok/mediator/cmd/cli/app"
	"github.com/stacklok/mediator/internal/util"
	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
)

var ruleType_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details for a rule type within a mediator control plane",
	Long: `The medic rule_type get subcommand lets you retrieve details for a rule type within a
mediator control plane.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			fmt.Fprintf(os.Stderr, "Error binding flags: %s\n", err)
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := viper.GetString("provider")
		format := viper.GetString("output")

		if format != app.JSON && format != app.YAML {
			return fmt.Errorf("error: invalid format: %s", format)
		}

		conn, err := util.GrpcForCommand(cmd)
		util.ExitNicelyOnError(err, "Error getting grpc connection")
		defer conn.Close()

		client := pb.NewPolicyServiceClient(conn)
		ctx, cancel := util.GetAppContext()
		defer cancel()

		id := viper.GetInt32("id")

		rtype, err := client.GetRuleTypeById(ctx, &pb.GetRuleTypeByIdRequest{
			Context: &pb.Context{
				Provider: provider,
				// TODO set up group if specified
				// Currently it's inferred from the authorization token
			},
			Id: id,
		})
		if err != nil {
			return fmt.Errorf("error getting rule type: %w", err)
		}

		if format == app.YAML {
			out, err := util.GetYamlFromProto(rtype)
			util.ExitNicelyOnError(err, "Error getting json from proto")
			fmt.Println(out)
		} else {
			out, err := util.GetJsonFromProto(rtype)
			util.ExitNicelyOnError(err, "Error getting json from proto")
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	ruleTypeCmd.AddCommand(ruleType_getCmd)
	ruleType_getCmd.Flags().Int32P("id", "i", 0, "ID for the policy to query")
	ruleType_getCmd.Flags().StringP("output", "o", app.YAML, "Output format (json or yaml)")
	ruleType_getCmd.Flags().StringP("provider", "p", "github", "Provider for the policy")
	// TODO set up group if specified

	if err := ruleType_getCmd.MarkFlagRequired("id"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %s\n", err)
		os.Exit(1)
	}

}
