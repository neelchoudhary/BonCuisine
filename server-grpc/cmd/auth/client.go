package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Auth Client")

	tls := true
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := user.NewUserServiceClient(conn)
	// getSavedRecipesTest(c)
	// removeSavedRecipeTest(c)
	// getSavedRecipesTest(c)
	// signupTest(c)
	loginTest(c)
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

func loginTest(c user.UserServiceClient) {
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
