package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neelchoudhary/boncuisine/api/models"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	"github.com/neelchoudhary/boncuisine/pkg/v1/user/repository"
)

type userServiceServer struct {
	userRecipeRepo *repository.UserRecipeRepository
	userRepo       *repository.UserRepository
}

// NewUserServiceServer TODO
func NewUserServiceServer(db *sql.DB) user.UserServiceServer {
	return &userServiceServer{userRecipeRepo: repository.NewUserRecipeRepository(db), userRepo: repository.NewUserRepository(db)}
}

func (s *userServiceServer) GetSavedRecipes(ctx context.Context, req *user.GetSavedRecipiesRequest) (*user.GetSavedRecipiesResponse, error) {
	fmt.Printf("Get saved recipes was invoked.\n")
	userID := req.GetUserId()
	recipes := s.userRecipeRepo.GetUserRecipes(userID)
	var pbSavedRecipes []*user.SavedRecipe
	for _, recipe := range recipes {
		pbSavedRecipes = append(pbSavedRecipes, dataToSavedRecipePb(recipe))
	}
	res := &user.GetSavedRecipiesResponse{
		SavedRecipes: pbSavedRecipes,
	}
	return res, nil
}

func (s *userServiceServer) AddSavedRecipe(ctx context.Context, req *user.AddSavedRecipeRequest) (*user.AddSavedRecipeResponse, error) {
	fmt.Printf("Add saved recipe was invoked.")
	userID := req.GetUserId()
	recipeID := req.GetRecipeId()
	s.userRecipeRepo.AddUserRecipe(userID, recipeID)
	res := &user.AddSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func (s *userServiceServer) RemoveSavedRecipe(ctx context.Context, req *user.RemoveSavedRecipeRequest) (*user.RemoveSavedRecipeResponse, error) {
	fmt.Printf("Remove saved recipe was invoked.")
	userID := req.GetUserId()
	recipeID := req.GetRecipeId()
	s.userRecipeRepo.RemoveUserRecipe(userID, recipeID)
	res := &user.RemoveSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func dataToSavedRecipePb(data models.SavedRecipe) *user.SavedRecipe {
	return &user.SavedRecipe{
		UserId:   data.UserID,
		RecipeId: data.RecipeID,
	}
}
