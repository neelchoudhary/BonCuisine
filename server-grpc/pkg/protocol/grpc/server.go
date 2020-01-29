package grpc

import (
	"context"
	"log"
	"net"

	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// RunServer registers gRPC service and run server
func RunServer(ctx context.Context, recipeServiceServer recipe.RecipeServiceServer, userServiceServer user.UserServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// TLS setup
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)

	// register service
	recipe.RegisterRecipeServiceServer(server, recipeServiceServer)
	user.RegisterUserServiceServer(server, userServiceServer)

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
