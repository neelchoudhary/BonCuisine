package main

import (
	"context"
	"fmt"
	"os"

	"github.com/neelchoudhary/boncuisine/api/driver"
	"github.com/neelchoudhary/boncuisine/pkg/protocol/grpc"
	account "github.com/neelchoudhary/boncuisine/pkg/v1/account/service"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/service"
)

func main() {
	env := "local"
	if err := grpc.RunServer(context.Background(), recipe.NewRecipeServiceServer(driver.ConnectDB(env)), account.NewAccountServiceServer(driver.ConnectDB(env)), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
