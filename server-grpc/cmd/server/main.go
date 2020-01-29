package main

import (
	"context"
	"fmt"
	"os"

	"github.com/neelchoudhary/boncuisine/api/driver"
	"github.com/neelchoudhary/boncuisine/pkg/protocol/grpc"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/service"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/service"
)

func main() {
	env := "local"
	if err := grpc.RunServer(context.Background(), recipe.NewRecipeServiceServer(driver.ConnectDB(env)), user.NewUserServiceServer(driver.ConnectDB(env)), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
