package main

import (
	"context"
	"fmt"
	"log"

	account "github.com/neelchoudhary/boncuisine/pkg/v1/account/api"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	//	c := recipe.NewRecipeServiceClient(conn)
	a := account.NewAccountServiceClient(conn)
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

func getSavedRecipesTest(c account.AccountServiceClient) {
	fmt.Println("Starting to do a GetSavedRecipes RPC...")
	req := &account.GetSavedRecipiesRequest{
		UserId: 1,
	}
	res, err := c.GetSavedRecipes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSavedRecipes())
}

func addSavedRecipeTest(c account.AccountServiceClient) {
	fmt.Println("Starting to do a AddSavedRecipe RPC...")
	req := &account.AddSavedRecipeRequest{
		UserId:   1,
		RecipeId: 4,
	}
	res, err := c.AddSavedRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}

func removeSavedRecipeTest(c account.AccountServiceClient) {
	fmt.Println("Starting to do a RemoveSavedRecipe RPC...")
	req := &account.RemoveSavedRecipeRequest{
		UserId:   1,
		RecipeId: 4,
	}
	res, err := c.RemoveSavedRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}
