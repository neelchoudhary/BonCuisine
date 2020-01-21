package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/neelchoudhary/boncuisine/models"
	"github.com/neelchoudhary/boncuisine/repositories"

	"github.com/gorilla/mux"
)

var userRecipes []models.SavedRecipe

func (c Controller) GetUserRecipes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRecipe models.SavedRecipe
		params := mux.Vars(r)

		userRecipes = []models.SavedRecipe{}
		savedRecipeRepo := repositories.SavedRecipeRepository{}

		id, err := strconv.Atoi(params["user_id"])
		logFatal(err)

		userRecipes = savedRecipeRepo.GetUserRecipes(db, userRecipe, userRecipes, id)

		json.NewEncoder(w).Encode(userRecipes)
	}
}

func (c Controller) AddUserRecipes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		userRecipes = []models.SavedRecipe{}

		user_id, err := strconv.Atoi(params["user_id"])
		logFatal(err)

		recipe_id, err := strconv.Atoi(params["recipe_id"])
		logFatal(err)

		savedRecipeRepo := repositories.SavedRecipeRepository{}
		user_id = savedRecipeRepo.AddUserRecipes(db, user_id, recipe_id)

		json.NewEncoder(w).Encode(user_id)
	}
}

func (c Controller) RemoveUserRecipes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		userRecipes = []models.SavedRecipe{}

		user_id, err := strconv.Atoi(params["user_id"])
		logFatal(err)

		recipe_id, err := strconv.Atoi(params["recipe_id"])
		logFatal(err)

		savedRecipeRepo := repositories.SavedRecipeRepository{}
		rowsDeleted := savedRecipeRepo.RemoveUserRecipes(db, user_id, recipe_id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
