package repositories

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/api/models"
)

type SavedRecipeRepository struct{}

func (sr SavedRecipeRepository) GetUserRecipes(db *sql.DB, userRecipe models.SavedRecipe, userRecipes []models.SavedRecipe, id int) []models.SavedRecipe {
	rows, err := db.Query("SELECT s.user_id, s.recipe_id FROM saved_recipes s INNER JOIN recipes r ON r.recipe_id = s.recipe_id AND s.user_id = $1;", id)
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userRecipe.UserID, &userRecipe.RecipeID)
		logFatal(err)

		userRecipes = append(userRecipes, userRecipe)
	}

	return userRecipes
}

func (sr SavedRecipeRepository) AddUserRecipes(db *sql.DB, user_id int, recipe_id int) int {
	// err :=
	db.QueryRow("INSERT INTO saved_recipes (user_id, recipe_id) VALUES ($1, $2);", user_id, recipe_id)

	// logFatal(err)

	return user_id
}

func (sr SavedRecipeRepository) RemoveUserRecipes(db *sql.DB, user_id int, recipe_id int) int64 {
	result, err := db.Exec("DELETE FROM saved_recipes WHERE saved_recipes.user_id = $1 AND saved_recipes.recipe_id = $2;", user_id, recipe_id)

	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
