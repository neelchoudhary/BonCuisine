package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/neelchoudhary/boncuisine/db/driver"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	recipeService "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/service"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	userService "github.com/neelchoudhary/boncuisine/pkg/v1/user/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port             = "8080"
	certFilePath     = "ssl/server.crt"
	keyFilePath      = "ssl/server.pem"
	secretServerCert = "/tls/server_cert"
	secretServerKey  = "/tls/server_key"
)

func main() {
	var env = flag.String("env", "local", "environment type: local, develop, staging, production")
	flag.Parse()

	// Write env to file
	err := utils.WriteEnvToFile(*env)
	utils.LogIfFatalAndExit(err, "Failed to write env to file")

	// Get secrets from AWS
	serverCert, err := utils.GetAWSSecret(*env + secretServerCert)
	utils.LogIfFatalAndExit(err)

	serverKey, err := utils.GetAWSSecret(*env + secretServerKey)
	utils.LogIfFatalAndExit(err)

	// Write secrets to file
	err = utils.WriteFile(certFilePath, serverCert)
	utils.LogIfFatalAndExit(err, "Failed to write server cert file:")

	err = utils.WriteFile(keyFilePath, serverKey)
	utils.LogIfFatalAndExit(err, "Failed to write server pem file:")

	db := driver.ConnectDB(*env)
	userService := userService.NewUserServiceServer(db)
	recipeService := recipeService.NewRecipeServiceServer(db)
	err = runServer(context.Background(), userService, recipeService)
	utils.LogIfFatalAndExit(err, "Failed to run server:")

}

// Starts the server
func runServer(ctx context.Context, userServiceServer user.UserServiceServer, recipeServiceServer recipe.RecipeServiceServer) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// TLS setup
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		creds, err := credentials.NewServerTLSFromFile(certFilePath, keyFilePath)
		utils.LogIfFatalAndExit(err, "Failed loading certificates:")
		opts = append(opts, grpc.Creds(creds))
	}

	// Add interceptor
	opts = append(opts, grpc.UnaryInterceptor(serverInterceptor))
	server := grpc.NewServer(opts...)

	// Register services
	user.RegisterUserServiceServer(server, userServiceServer)
	recipe.RegisterRecipeServiceServer(server, recipeServiceServer)

	// Start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

// Authorization unary interceptor function to handle authorize per RPC call
func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authorize when signing up for logging in
	if info.FullMethod != "/user.UserService/Signup" && info.FullMethod != "/user.UserService/Login" {
		userID, err := utils.AuthorizeToken(ctx)
		if err != nil {
			return nil, err
		}

		// Calls the handler with new context
		h, err := handler(utils.PassUserIDMetadata(ctx, userID), req)
		return h, err
	}

	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}
