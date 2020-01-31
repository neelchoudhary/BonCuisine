package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/db/models"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
)

// RecipeRepository struct
type RecipeRepository struct {
	db *sql.DB
}

// NewRecipeRepository sets the data source (e.g database)
func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

// GetAllRecipies Gets all recipes from db
func (r *RecipeRepository) GetAllRecipies() []models.Recipe {
	rows, err := r.db.Query("SELECT recipe_id, recipe_name, r_time, num_servings, difficulty, c.cuisine_name, i.image_data FROM recipes r INNER JOIN image_store i ON r.image_id = i.image_id INNER JOIN cuisines c ON r.cuisine_id = c.cuisine_id;")
	utils.LogFatal(err)

	defer rows.Close()
	recipes := make([]models.Recipe, 0)
	for rows.Next() {
		recipe := models.Recipe{}
		err := rows.Scan(&recipe.ID, &recipe.RecipeName, &recipe.Time, &recipe.Servings, &recipe.Difficulty, &recipe.CuisineName, &recipe.ImageData)
		utils.LogFatal(err)
		recipes = append(recipes, recipe)
	}

	return recipes
}

// GetRecipeIngredients Gets ingredients for a given recipe
func (r *RecipeRepository) GetRecipeIngredients(id int64) []models.RecipeIngredient {
	rows, err := r.db.Query("SELECT ing.ingredient_id, ing.ingredient_name, ri.amount, i.image_data FROM ingredients ing INNER JOIN recipe_ingredients ri ON ing.ingredient_id = ri.ingredient_id AND ri.recipe_id = $1 INNER JOIN image_store i ON ing.image_id = i.image_id;", id)
	utils.LogFatal(err)

	defer rows.Close()
	recipeIngredients := make([]models.RecipeIngredient, 0)
	for rows.Next() {
		recipeIngredient := models.RecipeIngredient{}
		err := rows.Scan(&recipeIngredient.IngredientID, &recipeIngredient.IngredientName, &recipeIngredient.Amount, &recipeIngredient.ImageData)
		utils.LogFatal(err)

		recipeIngredients = append(recipeIngredients, recipeIngredient)
	}

	return recipeIngredients
}
