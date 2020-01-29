package main

import (
	"context"
	"fmt"
	"log"

	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Client")

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

	//	c := recipe.NewRecipeServiceClient(conn)
	a := user.NewUserServiceClient(conn)
	//	getAllRecipesTest(c)
	//	getAllCuisinesTest(c)
	getSavedRecipesTest(a)
	removeSavedRecipeTest(a)
	getSavedRecipesTest(a)
}

func getAllRecipesTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetAllRecipes RPC...")
	req := &recipe.Empty{}
	res, err := c.GetAllRecipes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetRecipes())
}

func getAllCuisinesTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetAllCuisines RPC...")
	req := &recipe.Empty{}
	res, err := c.GetAllCuisines(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetCuisines())
}

func getSavedRecipesTest(c user.UserServiceClient) {
	fmt.Println("Starting to do a GetSavedRecipes RPC...")
	req := &user.GetSavedRecipiesRequest{
		UserId: 1,
	}
	res, err := c.GetSavedRecipes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSavedRecipes())
}

func addSavedRecipeTest(c user.UserServiceClient) {
	fmt.Println("Starting to do a AddSavedRecipe RPC...")
	req := &user.AddSavedRecipeRequest{
		UserId:   1,
		RecipeId: 4,
	}
	res, err := c.AddSavedRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}

func removeSavedRecipeTest(c user.UserServiceClient) {
	fmt.Println("Starting to do a RemoveSavedRecipe RPC...")
	req := &user.RemoveSavedRecipeRequest{
		UserId:   1,
		RecipeId: 4,
	}
	res, err := c.RemoveSavedRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}
