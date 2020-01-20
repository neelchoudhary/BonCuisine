package controllers

import (
	"boncuisine-mobile-app/Server2/models"
	"boncuisine-mobile-app/Server2/repositories"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

)

var cuisines []models.Cuisine

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetAllCuisines(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cuisine models.Cuisine
		cuisines = []models.Cuisine{}
		cuisineRepo := repositories.CuisineRepository{}

		cuisines = cuisineRepo.GetAllCuisines(db, cuisine, cuisines)

		json.NewEncoder(w).Encode(cuisines)
	}
}

