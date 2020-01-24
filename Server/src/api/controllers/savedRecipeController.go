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

var userRecipes []models.SavedRecipe

// GetUserRecipes godoc
// @Summary Get user saved recipes
// @Description Gets all recipes saved by a given user by user ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.SavedRecipe
// @Router /savedrecipes/{user_id}/ [get]
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

// AddUserRecipes godoc
// @Summary Add a saved recipe for a user
// @Description Adds a new saved recipe for a user given a user ID and recipe ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param recipe_id path int true "Recipe ID"
// @Success 200 {object} models.SavedRecipe
// @Router /savedrecipes/{user_id}/{recipe_id}/ [post]
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

// RemoveUserRecipes godoc
// @Summary Removes a saved recipe for a user
// @Description Removes an existing saved recipe for a user given a user ID and recipe ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param recipe_id path int true "Recipe ID"
// @Success 200 {object} models.SavedRecipe
// @Router /savedrecipes/{user_id}/{recipe_id}/ [delete]
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
