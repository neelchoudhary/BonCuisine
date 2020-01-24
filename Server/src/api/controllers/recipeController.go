package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/neelchoudhary/boncuisine/api/models"
	"github.com/neelchoudhary/boncuisine/api/repositories"

	"github.com/gorilla/mux"
)

var recipes []models.Recipe
var recipeIngredients []models.RecipeIngredient

// GetAllRecipes godoc
// @Summary Get all recipes
// @Description Gets all existing recipes.
// @Tags Recipes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Recipe
// @Router /recipes/ [get]
func (c Controller) GetAllRecipes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var recipe models.Recipe
		recipes = []models.Recipe{}
		recipeRepo := repositories.RecipeRepository{}

		recipes = recipeRepo.GetAllRecipies(db, recipe, recipes)

		json.NewEncoder(w).Encode(recipes)
	}
}

// GetRecipeIngredients godoc
// @Summary Get recipe ingredients
// @Description Gets all ingredients for a given recipe by recipe ID.
// @Tags Recipes
// @Accept  json
// @Produce  json
// @Param recipe_id path int true "Recipe ID"
// @Success 200 {array} models.RecipeIngredient
// @Router /recipe/{recipe_id}/ingredients/ [get]
func (c Controller) GetRecipeIngredients(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var recipeIngredient models.RecipeIngredient
		params := mux.Vars(r)

		recipeIngredients = []models.RecipeIngredient{}
		recipeRepo := repositories.RecipeRepository{}

		id, err := strconv.Atoi(params["recipe_id"])
		logFatal(err)

		recipeIngredients = recipeRepo.GetRecipeIngredients(db, recipeIngredient, recipeIngredients, id)

		json.NewEncoder(w).Encode(recipeIngredients)
	}
}
