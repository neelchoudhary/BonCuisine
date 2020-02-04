package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Auth Client")

	var port = flag.String("port", "8080", "server port")
	var certFilePath = flag.String("certFilePath", "ssl/ca.crt", "TLS ca cert file path")
	var accessTokenPath = flag.String("accessTokenPath", "cmd/auth/accessToken", "Access token path")

	tls := true
	opts := grpc.WithInsecure()
	if tls {
		// Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(*certFilePath, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial("localhost:"+*port, opts)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := user.NewUserServiceClient(conn)
	// signupTest(c)
	loginTest(c, *accessTokenPath)
}

func signupTest(c user.UserServiceClient) {
	fmt.Println("Starting to do a Signup RPC...")
	req := &user.SignupRequest{
		SignUpUser: &user.SignUpUser{
			Email:    "test@dev",
			Password: "test",
			Fullname: "Tester",
			Username: "Test Username",
		},
	}
	res, err := c.Signup(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}

func loginTest(c user.UserServiceClient, accessTokenPath string) {
	fmt.Println("Starting to do a Login RPC...")
	req := &user.LoginRequest{
		LoginUser: &user.LoginUser{
			Email:    "test@dev",
			Password: "test",
		},
	}
	res, err := c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
	log.Printf("Response from: %v", res.GetToken())
	err = ioutil.WriteFile("cmd/auth/accessToken", []byte(res.GetToken()), 0600)
	if err != nil {
		log.Fatal(err)
	}
}
