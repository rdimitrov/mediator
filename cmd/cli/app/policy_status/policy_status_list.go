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

package policy_status

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/stacklok/mediator/internal/util"
	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
)

var policystatus_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List policy status within a mediator control plane",
	Long: `The medic policy_status list subcommand lets you list policy status within a
mediator control plane for an specific provider/group or policy id.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			fmt.Fprintf(os.Stderr, "Error binding flags: %s\n", err)
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := util.GrpcForCommand(cmd)
		if err != nil {
			return fmt.Errorf("error getting grpc connection: %w", err)
		}
		defer conn.Close()

		client := pb.NewPolicyServiceClient(conn)
		ctx, cancel := util.GetAppContext()
		defer cancel()

		provider := viper.GetString("provider")
		group := viper.GetString("group")
		policyId := viper.GetInt32("policy")
		format := viper.GetString("output")
		all := viper.GetBool("all")

		if format != "json" && format != "yaml" {
			return fmt.Errorf("error: invalid format: %s", format)
		}

		if provider == "" {
			return fmt.Errorf("provider must be set")
		}

		if policyId == 0 {
			return fmt.Errorf("policy-id must be set")
		}

		req := &pb.GetPolicyStatusByIdRequest{
			Context: &pb.Context{
				Provider: provider,
			},
			PolicyId: policyId,
			EntitySelector: &pb.GetPolicyStatusByIdRequest_All{
				All: all,
			},
		}

		if group != "" {
			req.Context.Group = &group
		}

		resp, err := client.GetPolicyStatusById(ctx, req)
		if err != nil {
			return fmt.Errorf("error getting policy status: %w", err)
		}

		if format == "json" {
			out, err := util.GetJsonFromProto(resp)
			util.ExitNicelyOnError(err, "Error getting json from proto")
			fmt.Println(out)
		} else {
			out, err := util.GetYamlFromProto(resp)
			util.ExitNicelyOnError(err, "Error getting yaml from proto")
			fmt.Println(out)
		}

		return nil
	},
}

func init() {
	PolicyStatusCmd.AddCommand(policystatus_listCmd)
	policystatus_listCmd.Flags().StringP("provider", "p", "github", "Provider to list policy status for")
	policystatus_listCmd.Flags().StringP("group", "g", "", "group id to list policy status for")
	policystatus_listCmd.Flags().Int32P("policy", "i", 0, "policy id to list policy status for")
	policystatus_listCmd.Flags().StringP("output", "o", "yaml", "Output format (json or yaml)")
	policystatus_listCmd.Flags().BoolP("all", "a", false, "List all policy violations")
}
