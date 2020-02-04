package repository

import (
	"github.com/neelchoudhary/boncuisine/db/models"

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
func (r *CuisineRepository) GetAllCuisines() ([]models.Cuisine, error) {
	rows, err := r.db.Query("SELECT cuisine_id, cuisine_name FROM cuisines;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cuisines := make([]models.Cuisine, 0)
	for rows.Next() {
		cuisine := models.Cuisine{}
		err := rows.Scan(&cuisine.ID, &cuisine.CuisineName)
		if err != nil {
			return nil, err
		}
		cuisines = append(cuisines, cuisine)
	}

	return cuisines, nil
}
