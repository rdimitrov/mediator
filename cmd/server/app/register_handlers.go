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

package app

import (
	"context"
	"log"

	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/proto/v1"
	"github.com/stacklok/mediator/pkg/services"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func registerHandlers(ctx context.Context, gwmux *runtime.ServeMux, grpcAddress string, opts []grpc.DialOption) {
	// Register HealthService handler
	if err := pb.RegisterHealthServiceHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	// Register AuthUrlService handler
	if err := pb.RegisterAuthUrlServiceHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts); err != nil {
		log.Fatalf("failed to register gateway for AuthUrlService: %v", err)
	}

	// Register CallBackService handler
	if err := pb.RegisterCallBackServiceHandlerFromEndpoint(ctx, gwmux, grpcAddress, opts); err != nil {
		log.Fatalf("failed to register gateway for CallBackService: %v", err)
	}
}

func registerGRPCServices(s *grpc.Server) {
	// Register HealthService handler
	pb.RegisterHealthServiceServer(s, &services.Server{})

	// Register AuthUrlService handler
	pb.RegisterAuthUrlServiceServer(s, &services.Server{
		ClientID:     viper.GetString("github.client_id"),
		ClientSecret: viper.GetString("github.client_secret"),
	})
}