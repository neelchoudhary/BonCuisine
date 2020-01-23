package repositories

import (
	"github.com/neelchoudhary/boncuisine/api/models"

	"database/sql"

	"log"
)

type CuisineRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c CuisineRepository) GetAllCuisines(db *sql.DB, cuisine models.Cuisine, cuisines []models.Cuisine) []models.Cuisine {
	rows, err := db.Query("SELECT cuisine_id, cuisine_name FROM cuisines;")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cuisine.ID, &cuisine.CuisineName)
		logFatal(err)

		cuisines = append(cuisines, cuisine)
	}

	return cuisines
}
