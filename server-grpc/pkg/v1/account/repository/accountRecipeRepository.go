package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/api/models"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
)

// AccountRecipeRepository struct
type AccountRecipeRepository struct {
	db *sql.DB
}

// NewAccountRecipeRepository sets the data source (e.g database)
func NewAccountRecipeRepository(db *sql.DB) *AccountRecipeRepository {
	return &AccountRecipeRepository{db: db}
}

// GetAccountRecipes gets the saved recipes for this account
func (r *AccountRecipeRepository) GetAccountRecipes(userID int64) []models.SavedRecipe {
	rows, err := r.db.Query("SELECT s.user_id, s.recipe_id FROM saved_recipes s INNER JOIN recipes r ON r.recipe_id = s.recipe_id AND s.user_id = $1;", userID)
	utils.LogFatal(err)

	defer rows.Close()

	accountRecipes := make([]models.SavedRecipe, 0)
	for rows.Next() {
		accountRecipe := models.SavedRecipe{}
		err := rows.Scan(&accountRecipe.UserID, &accountRecipe.RecipeID)
		utils.LogFatal(err)
		accountRecipes = append(accountRecipes, accountRecipe)
	}

	return accountRecipes
}

// AddAccountRecipe adds a new saved recipe for this account
func (r *AccountRecipeRepository) AddAccountRecipe(userID int64, recipeID int64) int64 {
	r.db.QueryRow("INSERT INTO saved_recipes (user_id, recipe_id) VALUES ($1, $2);", userID, recipeID)

	return userID
}

// RemoveAccountRecipe Removes a saved recipe from this account
func (r *AccountRecipeRepository) RemoveAccountRecipe(userID int64, recipeID int64) int64 {
	result, err := r.db.Exec("DELETE FROM saved_recipes WHERE saved_recipes.user_id = $1 AND saved_recipes.recipe_id = $2;", userID, recipeID)

	utils.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	utils.LogFatal(err)

	return rowsDeleted
}
