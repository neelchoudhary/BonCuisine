package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/api/models"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
)

// UserRecipeRepository struct
type UserRecipeRepository struct {
	db *sql.DB
}

// NewUserRecipeRepository sets the data source (e.g database)
func NewUserRecipeRepository(db *sql.DB) *UserRecipeRepository {
	return &UserRecipeRepository{db: db}
}

// GetUserRecipes gets the saved recipes for this user
func (r *UserRecipeRepository) GetUserRecipes(userID string) []models.SavedRecipe {
	rows, err := r.db.Query("SELECT s.user_id, s.recipe_id FROM saved_recipes s INNER JOIN recipes r ON r.recipe_id = s.recipe_id AND s.user_id = $1;", userID)
	utils.LogFatal(err)

	defer rows.Close()

	userRecipes := make([]models.SavedRecipe, 0)
	for rows.Next() {
		userRecipe := models.SavedRecipe{}
		err := rows.Scan(&userRecipe.UserID, &userRecipe.RecipeID)
		utils.LogFatal(err)
		userRecipes = append(userRecipes, userRecipe)
	}

	return userRecipes
}

// AddUserRecipe adds a new saved recipe for this user
func (r *UserRecipeRepository) AddUserRecipe(userID string, recipeID int64) string {
	r.db.QueryRow("INSERT INTO saved_recipes (user_id, recipe_id) VALUES ($1, $2);", userID, recipeID)

	return userID
}

// RemoveUserRecipe Removes a saved recipe from this user
func (r *UserRecipeRepository) RemoveUserRecipe(userID string, recipeID int64) int64 {
	result, err := r.db.Exec("DELETE FROM saved_recipes WHERE saved_recipes.user_id = $1 AND saved_recipes.recipe_id = $2;", userID, recipeID)

	utils.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	utils.LogFatal(err)

	return rowsDeleted
}
