package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/neelchoudhary/boncuisine/api/driver"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	recipeService "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	env := "local"
	if err := runServer(context.Background(), recipeService.NewRecipeServiceServer(driver.ConnectDB(env)), "3001"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer registers gRPC service and run server
func runServer(ctx context.Context, recipeServiceServer recipe.RecipeServiceServer, port string) error {
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

	opts = append(opts, grpc.UnaryInterceptor(serverInterceptor))
	server := grpc.NewServer(opts...)

	// register service
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
	// Skip authorize when GetJWT is requested
	fmt.Println(info.FullMethod)
	// TODO change
	if info.FullMethod != "/user.UserService/Signup" && info.FullMethod != "/user.UserService/Login" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]

	const prefix = "Bearer "
	if !strings.HasPrefix(token, prefix) {
		return status.Error(codes.Unauthenticated, `missing "Bearer " prefix in "Authorization" header`)
	}

	// if strings.TrimPrefix(token, prefix) != a.Token {
	// 	return status.Error(codes.Unauthenticated, "invalid token")
	// }

	token = strings.TrimPrefix(token, prefix)
	// validateToken function validates the token
	err := validateToken(token)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	return nil
}

func validateToken(tokenString string) error {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("verySecretSecret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"], claims["email"], claims["name"])
	} else {
		fmt.Println(err)
		return err
	}
	return nil
}
