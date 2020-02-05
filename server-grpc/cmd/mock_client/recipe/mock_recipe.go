package recipe

import (
	"context"
	"fmt"
	"log"

	"github.com/neelchoudhary/boncuisine/pkg/utils"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	"google.golang.org/grpc"
)

// DialServer connects to server with JWT Credentials
func DialServer(address string, opts grpc.DialOption, accessTokenPath string) (*grpc.ClientConn, error) {
	// Get token from safe space
	data, err := utils.ReadFile(accessTokenPath)
	utils.LogIfFatalAndExit(err, "Unable to read access token file:")

	if string(data) == "" {
		log.Fatal("No token!")
	}
	jwtCreds := utils.GetTokenAuth(string(data))
	utils.LogIfFatalAndExit(err)
	return grpc.Dial(address, opts, grpc.WithPerRPCCredentials(jwtCreds))
}

// RunTests runs client tests
func RunTests(conn *grpc.ClientConn) {
	c := recipe.NewRecipeServiceClient(conn)
	// getAllRecipesTest(c)
	// getAllCuisinesTest(c)
	// getRecipeIngredients(c)
	getSavedRecipesTest(c)
	addSavedRecipeTest(c)
	getSavedRecipesTest(c)
	removeSavedRecipeTest(c)
	getSavedRecipesTest(c)
}

func getAllRecipesTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetAllRecipes RPC...")
	req := &recipe.Empty{}
	res, err := c.GetAllRecipes(context.Background(), req)
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetRecipes())
}

func getAllCuisinesTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetAllCuisines RPC...")
	req := &recipe.Empty{}
	res, err := c.GetAllCuisines(context.Background(), req)
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetCuisines())
}

func getRecipeIngredients(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetRecipeIngredients RPC...")
	req := &recipe.GetRecipeIngredientsRequest{
		RecipeId: 1,
	}
	res, err := c.GetRecipeIngredients(context.Background(), req)
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetRecipeIngredients())
}

func getSavedRecipesTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a GetSavedRecipes RPC...")
	req := &recipe.Empty{}
	res, err := c.GetSavedRecipes(context.Background(), req)
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetSavedRecipes())
}

func addSavedRecipeTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a AddSavedRecipe RPC...")
	req := &recipe.AddSavedRecipeRequest{
		RecipeId: 4,
	}
	res, err := c.AddSavedRecipe(context.Background(), req)
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetSuccess())
}

func removeSavedRecipeTest(c recipe.RecipeServiceClient) {
	fmt.Println("Starting to do a RemoveSavedRecipe RPC...")
	req := &recipe.RemoveSavedRecipeRequest{
		RecipeId: 4,
	}
	res, err := c.RemoveSavedRecipe(context.Background(), req)
	utils.LogIfFatalAndExit(err, "Error while calling RPC:")
	log.Printf("Response from: %v", res.GetSuccess())
}
