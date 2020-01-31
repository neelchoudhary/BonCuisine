package recipe

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neelchoudhary/boncuisine/db/models"
	repository "github.com/neelchoudhary/boncuisine/db/repositories"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
)

type recipeServiceServer struct {
	recipeRepo  *repository.RecipeRepository
	cuisineRepo *repository.CuisineRepository
}

// NewRecipeServiceServer creates Chat service object
func NewRecipeServiceServer(db *sql.DB) recipe.RecipeServiceServer {
	return &recipeServiceServer{recipeRepo: repository.NewRecipeRepository(db), cuisineRepo: repository.NewCuisineRepository(db)}
}

func (s *recipeServiceServer) GetAllRecipes(ctx context.Context, req *recipe.Empty) (*recipe.GetAllRecipeResponse, error) {
	fmt.Printf("Get all recipes was invoked.\n")
	recipes := s.recipeRepo.GetAllRecipies()
	fmt.Printf("s1\n")
	var pbRecipes []*recipe.Recipe
	for _, recipe := range recipes {
		pbRecipes = append(pbRecipes, dataToRecipePb(recipe))
	}
	fmt.Printf("s2\n")

	res := &recipe.GetAllRecipeResponse{
		Recipes: pbRecipes,
	}
	fmt.Printf("s3\n")
	return res, nil
}

func (s *recipeServiceServer) GetAllCuisines(ctx context.Context, req *recipe.Empty) (*recipe.GetAllCuisinesResponse, error) {
	fmt.Printf("Get all cuisines was invoked.")
	cuisines := s.cuisineRepo.GetAllCuisines()

	var pbCuisines []*recipe.Cuisine
	for _, cuisine := range cuisines {
		pbCuisines = append(pbCuisines, dataToCuisinePb(cuisine))
	}

	res := &recipe.GetAllCuisinesResponse{
		Cuisines: pbCuisines,
	}
	return res, nil
}

func (s *recipeServiceServer) GetRecipeIngredients(ctx context.Context, req *recipe.GetRecipeIngredientsRequest) (*recipe.GetRecipeIngredientsResponse, error) {
	fmt.Printf("Get recipe ingredients was invoked.")
	recipeID := req.GetRecipeId()
	recipeIngredients := s.recipeRepo.GetRecipeIngredients(recipeID)

	var pbRecipeIngredients []*recipe.RecipeIngredient
	for _, recipeIngredient := range recipeIngredients {
		pbRecipeIngredients = append(pbRecipeIngredients, dataToRecipeIngredientPb(recipeIngredient))
	}

	res := &recipe.GetRecipeIngredientsResponse{
		RecipeIngredients: pbRecipeIngredients,
	}
	return res, nil
}

func dataToRecipePb(data models.Recipe) *recipe.Recipe {
	return &recipe.Recipe{
		Id:         data.ID,
		Name:       data.RecipeName,
		Time:       data.Time,
		Servings:   data.Servings,
		Difficulty: data.Difficulty,
		Cuisine:    data.CuisineName,
		Image:      data.ImageData,
	}
}

func dataToCuisinePb(data models.Cuisine) *recipe.Cuisine {
	return &recipe.Cuisine{
		Id:   data.ID,
		Name: data.CuisineName,
	}
}

func dataToRecipeIngredientPb(data models.RecipeIngredient) *recipe.RecipeIngredient {
	return &recipe.RecipeIngredient{
		Id:     data.IngredientID,
		Name:   data.IngredientName,
		Amount: data.Amount,
		Image:  data.ImageData,
	}
}
