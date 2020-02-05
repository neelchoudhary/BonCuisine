package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/neelchoudhary/boncuisine/cmd/mock_client/auth"
	"github.com/neelchoudhary/boncuisine/cmd/mock_client/recipe"
	"github.com/neelchoudhary/boncuisine/pkg/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port             = "8080"
	certFilePath     = "ssl/ca.crt"
	accessTokenPath  = "cmd/mock_client/auth/accessToken"
	secretClientCert = "/tls/client_cert"
)

func main() {
	fmt.Println("Starting Client..")

	// Flags
	var env = flag.String("env", "local", "environment type: local, develop, staging, production")
	var client = flag.String("c", "recipe", "The client to run: auth, recipe")
	flag.Parse()

	// TLS
	clientCert, err := utils.GetAWSSecret(*env + secretClientCert)
	utils.LogIfFatalAndExit(err)
	err = ioutil.WriteFile(certFilePath, []byte(clientCert), 0600)
	utils.LogIfFatalAndExit(err, "Failed to write client cert file")
	creds, err := credentials.NewClientTLSFromFile(certFilePath, "")
	utils.LogIfFatalAndExit(err, "Error while loading CA trust certificate:")
	opts := grpc.WithTransportCredentials(creds)

	var domain string
	if *env == "local" {
		domain = "localhost:"
	} else if *env == "develop" {
		domain = "dev.boncuisine-server.com:"
	}

	var conn *grpc.ClientConn
	address := domain + port
	if *client == "auth" {
		conn, err = auth.DialServer(address, opts)
		fmt.Println("Started Auth Client")
	} else if *client == "recipe" {
		conn, err = recipe.DialServer(address, opts, accessTokenPath)
		fmt.Println("Started Recipe Client")
	} else {
		log.Fatal("Failed to start client: unknown client, use auth or recipe")
	}

	utils.LogIfFatalAndExit(err, "Failed to connect:")
	defer conn.Close()

	if *client == "auth" {
		auth.RunTests(conn, accessTokenPath)
	} else if *client == "recipe" {
		recipe.RunTests(conn)
	}
}
