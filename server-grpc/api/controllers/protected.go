package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Protected godoc
// @Summary TODO
// @Description TODO.
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Router /protectedEndpoint/ [get]
func (c Controller) Protected(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("success.")
	}
}
