package recipe

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neelchoudhary/boncuisine/pkg/utils"

	"google.golang.org/grpc/codes"

	"github.com/neelchoudhary/boncuisine/db/models"
	repository "github.com/neelchoudhary/boncuisine/db/repositories"
	recipe "github.com/neelchoudhary/boncuisine/pkg/v1/recipe/api"
	"google.golang.org/grpc/status"
)

type recipeServiceServer struct {
	recipeRepo      *repository.RecipeRepository
	cuisineRepo     *repository.CuisineRepository
	savedRecipeRepo *repository.SavedRecipeRepository
}

// NewRecipeServiceServer creates recipe service object
func NewRecipeServiceServer(db *sql.DB) recipe.RecipeServiceServer {
	return &recipeServiceServer{recipeRepo: repository.NewRecipeRepository(db), cuisineRepo: repository.NewCuisineRepository(db), savedRecipeRepo: repository.NewSavedRecipeRepository(db)}
}

func (s *recipeServiceServer) GetAllRecipes(ctx context.Context, req *recipe.Empty) (*recipe.GetAllRecipeResponse, error) {
	recipes, err := s.recipeRepo.GetAllRecipies()
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting all recipes: %s", err.Error()))
	}
	var pbRecipes []*recipe.Recipe
	for _, recipe := range recipes {
		pbRecipes = append(pbRecipes, dataToRecipePb(recipe))
	}

	res := &recipe.GetAllRecipeResponse{
		Recipes: pbRecipes,
	}
	return res, nil
}

func (s *recipeServiceServer) GetAllCuisines(ctx context.Context, req *recipe.Empty) (*recipe.GetAllCuisinesResponse, error) {
	cuisines, err := s.cuisineRepo.GetAllCuisines()
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting all cuisines: %s", err.Error()))
	}

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
	recipeID := req.GetRecipeId()
	recipeIngredients, err := s.recipeRepo.GetRecipeIngredients(recipeID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error getting recipe ingredients: %s", err.Error()))
	}

	var pbRecipeIngredients []*recipe.RecipeIngredient
	for _, recipeIngredient := range recipeIngredients {
		pbRecipeIngredients = append(pbRecipeIngredients, dataToRecipeIngredientPb(recipeIngredient))
	}

	res := &recipe.GetRecipeIngredientsResponse{
		RecipeIngredients: pbRecipeIngredients,
	}
	return res, nil
}

func (s *recipeServiceServer) GetSavedRecipes(ctx context.Context, req *recipe.Empty) (*recipe.GetSavedRecipiesResponse, error) {
	// Get UserID from metadata
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	recipes, err := s.savedRecipeRepo.GetSavedRecipes(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting saved recipes: %s", err.Error()))
	}
	var pbSavedRecipes []*recipe.SavedRecipe
	for _, recipe := range recipes {
		pbSavedRecipes = append(pbSavedRecipes, dataToSavedRecipePb(recipe))
	}
	res := &recipe.GetSavedRecipiesResponse{
		SavedRecipes: pbSavedRecipes,
	}
	return res, nil
}

func (s *recipeServiceServer) AddSavedRecipe(ctx context.Context, req *recipe.AddSavedRecipeRequest) (*recipe.AddSavedRecipeResponse, error) {
	// Get UserID from metadata
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	recipeID := req.GetRecipeId()
	err = s.savedRecipeRepo.AddSavedRecipe(userID, recipeID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error adding saved recipe: %s", err.Error()))
	}
	res := &recipe.AddSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func (s *recipeServiceServer) RemoveSavedRecipe(ctx context.Context, req *recipe.RemoveSavedRecipeRequest) (*recipe.RemoveSavedRecipeResponse, error) {
	// Get UserID from metadata
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	recipeID := req.GetRecipeId()
	err = s.savedRecipeRepo.RemoveSavedRecipe(userID, recipeID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error removing saved recipe: %s", err.Error()))
	}
	res := &recipe.RemoveSavedRecipeResponse{
		Success: true,
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

func dataToSavedRecipePb(data models.SavedRecipe) *recipe.SavedRecipe {
	return &recipe.SavedRecipe{
		UserId:   data.UserID,
		RecipeId: data.RecipeID,
	}
}
