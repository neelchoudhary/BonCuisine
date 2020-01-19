package router

import (
	"dormeter.com/middleware"
	"github.com/gorilla/mux"
)

//Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.AddUser).Methods("POST", "OPTIONS")

	return router
}
