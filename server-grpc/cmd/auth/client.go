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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

const (
	cert = "client_cert"
)

func getTLSKeys(env string, tlsType string) string {
	var secretName string
	secretName = env + "/tls/" + tlsType
	region := "us-east-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		log.Fatal(err.(awserr.Error))
	}

	return *result.SecretString
}

func main() {
	fmt.Println("Auth Client")
	var env = flag.String("env", "local", "environment type: local, develop, staging, production")
	var port = flag.String("port", "8080", "server port")
	var certFilePath = flag.String("certFilePath", "ssl/ca.crt", "TLS ca cert file path")
	var accessTokenPath = flag.String("accessTokenPath", "cmd/auth/accessToken", "Access token path")

	flag.Parse()

	tls := true
	opts := grpc.WithInsecure()
	if tls {
		err := ioutil.WriteFile(*certFilePath, []byte(getTLSKeys(*env, cert)), 0600)
		if err != nil {
			fmt.Println("Failed to write client cert file")
		}

		// Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(*certFilePath, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	var domain string
	if *env == "local" {
		domain = "localhost:"
	} else if *env == "develop" {
		domain = "dev.boncuisine-server.com:"
	}
	conn, err := grpc.Dial(domain+*port, opts)
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
