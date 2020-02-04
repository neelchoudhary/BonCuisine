package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/db/models"
)

// SavedRecipeRepository struct
type SavedRecipeRepository struct {
	db *sql.DB
}

// NewSavedRecipeRepository sets the data source (e.g database)
func NewSavedRecipeRepository(db *sql.DB) *SavedRecipeRepository {
	return &SavedRecipeRepository{db: db}
}

// GetSavedRecipes gets the saved recipes for this user
func (r *SavedRecipeRepository) GetSavedRecipes(userID string) ([]models.SavedRecipe, error) {
	rows, err := r.db.Query("SELECT s.user_id, s.recipe_id FROM saved_recipes s INNER JOIN recipes r ON r.recipe_id = s.recipe_id AND s.user_id = $1;", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userRecipes := make([]models.SavedRecipe, 0)
	for rows.Next() {
		userRecipe := models.SavedRecipe{}
		err := rows.Scan(&userRecipe.UserID, &userRecipe.RecipeID)
		if err != nil {
			return nil, err
		}
		userRecipes = append(userRecipes, userRecipe)
	}

	return userRecipes, nil
}

// AddSavedRecipe adds a new saved recipe for this user
func (r *SavedRecipeRepository) AddSavedRecipe(userID string, recipeID int64) error {
	_, err := r.db.Exec("INSERT INTO saved_recipes (user_id, recipe_id) VALUES ($1, $2);", userID, recipeID)

	if err != nil {
		return err
	}

	return nil
}

// RemoveSavedRecipe Removes a saved recipe from this user
func (r *SavedRecipeRepository) RemoveSavedRecipe(userID string, recipeID int64) error {
	_, err := r.db.Exec("DELETE FROM saved_recipes WHERE saved_recipes.user_id = $1 AND saved_recipes.recipe_id = $2;", userID, recipeID)

	if err != nil {
		return err
	}

	return nil
}
