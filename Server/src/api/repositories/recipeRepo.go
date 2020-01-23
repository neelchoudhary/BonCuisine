package repositories

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/api/models"
)

type RecipeRepository struct{}

func (r RecipeRepository) GetAllRecipies(db *sql.DB, recipe models.Recipe, recipes []models.Recipe) []models.Recipe {
	rows, err := db.Query("SELECT recipe_id, recipe_name, r_time, num_servings, difficulty, c.cuisine_name, i.image_data FROM recipes r INNER JOIN image_store i ON r.image_id = i.image_id INNER JOIN cuisines c ON r.cuisine_id = c.cuisine_id;")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&recipe.ID, &recipe.RecipeName, &recipe.Time, &recipe.Servings, &recipe.Difficulty, &recipe.CuisineName, &recipe.ImageData)
		logFatal(err)

		recipes = append(recipes, recipe)
	}

	return recipes
}

func (r RecipeRepository) GetRecipeIngredients(db *sql.DB, recipeIngredient models.RecipeIngredient, recipeIngredients []models.RecipeIngredient, id int) []models.RecipeIngredient {
	rows, err := db.Query("SELECT ing.ingredient_id, ing.ingredient_name, ri.amount, i.image_data FROM ingredients ing INNER JOIN recipe_ingredients ri ON ing.ingredient_id = ri.ingredient_id AND ri.recipe_id = $1 INNER JOIN image_store i ON ing.image_id = i.image_id;", id)
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&recipeIngredient.IngredientID, &recipeIngredient.IngredientName, &recipeIngredient.Amount, &recipeIngredient.ImageData)
		logFatal(err)

		recipeIngredients = append(recipeIngredients, recipeIngredient)
	}

	return recipeIngredients
}
