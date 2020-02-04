package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Recipe Client")

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

	// Get token from safe space
	data, err := ioutil.ReadFile(*accessTokenPath)
	if err != nil {
		log.Fatalf("Unable to read access token file: %v", err)
	}
	if string(data) == "" {
		log.Fatalf("No token! ")
	}
	jwtCreds := tokenAuth{string(data)}
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:"+*port, opts, grpc.WithPerRPCCredentials(jwtCreds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := recipe.NewRecipeServiceClient(conn)
	// getAllRecipesTest(c)
	//getAllCuisinesTest(c)
	// getRecipeIngredients(c)
	getSavedRecipesTest(c)
	addSavedRecipeTest(c)
	getSavedRecipesTest(c)
	removeSavedRecipeTest(c)
	getSavedRecipesTest(c)
}

type tokenAuth struct {
	token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
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

func getRecipeIngredients(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetRecipeIngredients RPC...")
	req := &recipe.GetRecipeIngredientsRequest{
		RecipeId: 1,
	}
	res, err := c.GetRecipeIngredients(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetRecipeIngredients())
}

func getSavedRecipesTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetSavedRecipes RPC...")
	req := &recipe.Empty{}
	res, err := c.GetSavedRecipes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSavedRecipes())
}

func addSavedRecipeTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a AddSavedRecipe RPC...")
	req := &recipe.AddSavedRecipeRequest{
		RecipeId: 4,
	}
	res, err := c.AddSavedRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}

func removeSavedRecipeTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a RemoveSavedRecipe RPC...")
	req := &recipe.RemoveSavedRecipeRequest{
		RecipeId: 4,
	}
	res, err := c.RemoveSavedRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling RPC: %v", err)
	}
	log.Printf("Response from: %v", res.GetSuccess())
}
