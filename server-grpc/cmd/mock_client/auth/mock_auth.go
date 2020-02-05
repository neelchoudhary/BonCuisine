package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/neelchoudhary/boncuisine/pkg/utils"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	"google.golang.org/grpc"
)

// DialServer connects to server
func DialServer(address string, opts grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(address, opts)
}

// RunTests runs client tests
func RunTests(conn *grpc.ClientConn, accessTokenPath string) {
	c := user.NewUserServiceClient(conn)
	// signupTest(c)
	loginTest(c, accessTokenPath)
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
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
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
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetSuccess())
	err = utils.WriteFile(accessTokenPath, res.GetToken())
	utils.LogIfFatalAndExit(err, "Error saving token to file:")
}
