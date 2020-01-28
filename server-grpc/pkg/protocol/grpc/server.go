package grpc

import (
	"context"
	"log"
	"net"

	account "github.com/neelchoudhary/boncuisine/pkg/v1/account/api"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"

	"google.golang.org/grpc"
)

// RunServer registers gRPC service and run server
func RunServer(ctx context.Context, recipeServiceServer recipe.RecipeServiceServer, accountServiceServer account.AccountServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	recipe.RegisterRecipeServiceServer(server, recipeServiceServer)
	account.RegisterAccountServiceServer(server, accountServiceServer)

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
