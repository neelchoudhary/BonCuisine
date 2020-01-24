package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/neelchoudhary/boncuisine/api/repositories"

	"github.com/neelchoudhary/boncuisine/api/models"
)

var cuisines []models.Cuisine

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllCuisines godoc
// @Summary Get all cuisines
// @Description Gets all existing cuisines.
// @Tags Cuisines
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Cuisine
// @Router /cuisines/ [get]
func (c Controller) GetAllCuisines(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cuisine models.Cuisine
		cuisines = []models.Cuisine{}
		cuisineRepo := repositories.CuisineRepository{}

		cuisines = cuisineRepo.GetAllCuisines(db, cuisine, cuisines)

		json.NewEncoder(w).Encode(cuisines)
	}
}
