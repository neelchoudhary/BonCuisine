package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"api.com/router"

	"api.com/models"
)

//I haven't completed refactoring yet(yall are gonna need to change it suit postgresql over mongodb)

//Documents This is a comment
var Documents []models.Document

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllDocuments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDocuments")
	json.NewEncoder(w).Encode(Documents)
}

func main() {

	r := router.Router()
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
