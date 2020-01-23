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

func (c Controller) GetAllRecipes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var recipe models.Recipe
		recipes = []models.Recipe{}
		recipeRepo := repositories.RecipeRepository{}

		recipes = recipeRepo.GetAllRecipies(db, recipe, recipes)

		json.NewEncoder(w).Encode(recipes)
	}
}

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
