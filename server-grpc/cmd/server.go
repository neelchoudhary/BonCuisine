package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/neelchoudhary/boncuisine/db/driver"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	recipeService "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/service"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	userService "github.com/neelchoudhary/boncuisine/pkg/v1/user/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// DB ...
type DB struct {
	db *sql.DB
}

const (
	cert = "server_cert"
	key  = "server_key"
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
	var env = flag.String("env", "local", "environment type: local, develop, staging, production")
	var port = flag.String("port", "8080", "server port")
	var certFilePath = flag.String("certFilePath", "ssl/server.crt", "TLS cert file path")
	var keyFilePath = flag.String("keyFilePath", "ssl/server.pem", "TLS key file path")
	//	var jwtSecret = flag.String("jwtSecret", "", "JWT secret")

	err := ioutil.WriteFile(*certFilePath, []byte(getTLSKeys(*env, cert)), 0600)
	if err != nil {
		fmt.Println("Failed to write server cert file")
	}
	err = ioutil.WriteFile(*keyFilePath, []byte(getTLSKeys(*env, key)), 0600)
	if err != nil {
		fmt.Println("Failed to write server pem file")
	}

	flag.Parse()

	db := DB{db: driver.ConnectDB(*env)}.db
	userService := userService.NewUserServiceServer(db)
	recipeService := recipeService.NewRecipeServiceServer(db)
	if err := runServer(context.Background(), userService, recipeService, *port, *certFilePath, *keyFilePath); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer registers gRPC service and run server
func runServer(ctx context.Context, userServiceServer user.UserServiceServer, recipeServiceServer recipe.RecipeServiceServer, port string, certFilePath string, keyFilePath string) error {
	listen, err := net.Listen("tcp", "dev.boncuisine-server.com:"+port)
	if err != nil {
		return err
	}

	// TLS setup
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		creds, sslErr := credentials.NewServerTLSFromFile(certFilePath, keyFilePath)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	opts = append(opts, grpc.UnaryInterceptor(serverInterceptor))
	server := grpc.NewServer(opts...)

	// register services
	user.RegisterUserServiceServer(server, userServiceServer)
	recipe.RegisterRecipeServiceServer(server, recipeServiceServer)

	// start gRPC server
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
		userID, err := authorize(ctx)
		if err != nil {
			return nil, err
		}
		md := metadata.Pairs("userid", userID)
		ctx = metadata.NewIncomingContext(ctx, md)
		// Calls the handler
		h, err := handler(ctx, req)
		return h, err
	}

	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]

	const prefix = "Bearer "
	if !strings.HasPrefix(token, prefix) {
		return "", status.Error(codes.Unauthenticated, `missing "Bearer " prefix in "Authorization" header`)
	}

	// if strings.TrimPrefix(token, prefix) != a.Token {
	// 	return status.Error(codes.Unauthenticated, "invalid token")
	// }

	token = strings.TrimPrefix(token, prefix)
	// validateToken function validates the token
	userID, err := validateToken(token)

	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, err.Error())
	}
	return userID, nil
}

func validateToken(tokenString string) (string, error) {

	claims := &userService.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("verySecretSecret"), nil
	})
	if err != nil || !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
