package account

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neelchoudhary/boncuisine/api/models"
	account "github.com/neelchoudhary/boncuisine/pkg/v1/account/api"
	"github.com/neelchoudhary/boncuisine/pkg/v1/account/repository"
)

type accountServiceServer struct {
	accountRecipeRepo *repository.AccountRecipeRepository
	accountRepo       *repository.AccountRepository
}

// NewAccountServiceServer TODO
func NewAccountServiceServer(db *sql.DB) account.AccountServiceServer {
	return &accountServiceServer{accountRecipeRepo: repository.NewAccountRecipeRepository(db), accountRepo: repository.NewAccountRepository(db)}
}

func (s *accountServiceServer) GetSavedRecipes(ctx context.Context, req *account.GetSavedRecipiesRequest) (*account.GetSavedRecipiesResponse, error) {
	fmt.Printf("Get saved recipes was invoked.\n")
	userID := req.GetUserId()
	recipes := s.accountRecipeRepo.GetAccountRecipes(userID)
	var pbSavedRecipes []*account.SavedRecipe
	for _, recipe := range recipes {
		pbSavedRecipes = append(pbSavedRecipes, dataToSavedRecipePb(recipe))
	}
	res := &account.GetSavedRecipiesResponse{
		SavedRecipes: pbSavedRecipes,
	}
	return res, nil
}

func (s *accountServiceServer) AddSavedRecipe(ctx context.Context, req *account.AddSavedRecipeRequest) (*account.AddSavedRecipeResponse, error) {
	fmt.Printf("Add saved recipe was invoked.")
	userID := req.GetUserId()
	recipeID := req.GetRecipeId()
	s.accountRecipeRepo.AddAccountRecipe(userID, recipeID)
	res := &account.AddSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func (s *accountServiceServer) RemoveSavedRecipe(ctx context.Context, req *account.RemoveSavedRecipeRequest) (*account.RemoveSavedRecipeResponse, error) {
	fmt.Printf("Remove saved recipe was invoked.")
	userID := req.GetUserId()
	recipeID := req.GetRecipeId()
	s.accountRecipeRepo.RemoveAccountRecipe(userID, recipeID)
	res := &account.RemoveSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func dataToSavedRecipePb(data models.SavedRecipe) *account.SavedRecipe {
	return &account.SavedRecipe{
		UserId:   data.UserID,
		RecipeId: data.RecipeID,
	}
}
