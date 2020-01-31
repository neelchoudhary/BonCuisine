package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/neelchoudhary/boncuisine/api/driver"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	userService "github.com/neelchoudhary/boncuisine/pkg/v1/user/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	env := "local"
	if err := runServer(context.Background(), userService.NewUserServiceServer(driver.ConnectDB(env)), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer registers gRPC service and run server
func runServer(ctx context.Context, userServiceServer user.UserServiceServer, port string) error {
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

	// register auth service
	user.RegisterUserServiceServer(server, userServiceServer)

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
