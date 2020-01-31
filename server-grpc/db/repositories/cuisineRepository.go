package repository

import (
	"github.com/neelchoudhary/boncuisine/db/models"
	"github.com/neelchoudhary/boncuisine/pkg/utils"

	"database/sql"
)

// CuisineRepository struct
type CuisineRepository struct {
	db *sql.DB
}

// NewCuisineRepository sets the data source (e.g database)
func NewCuisineRepository(db *sql.DB) *CuisineRepository {
	return &CuisineRepository{db: db}
}

// GetAllCuisines Gets all cuisines from db
func (r *CuisineRepository) GetAllCuisines() []models.Cuisine {
	rows, err := r.db.Query("SELECT cuisine_id, cuisine_name FROM cuisines;")
	utils.LogFatal(err)

	defer rows.Close()

	cuisines := make([]models.Cuisine, 0)
	for rows.Next() {
		cuisine := models.Cuisine{}
		err := rows.Scan(&cuisine.ID, &cuisine.CuisineName)
		utils.LogFatal(err)
		cuisines = append(cuisines, cuisine)
	}

	return cuisines
}
