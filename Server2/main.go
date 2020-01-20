package main

import (
	"boncuisine-mobile-app/Server2/driver"
	"boncuisine-mobile-app/Server2/controllers"
	"boncuisine-mobile-app/Server2/utils"
	"database/sql"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	// "strconv"
)

var db *sql.DB


func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}


func main() {
	db = driver.ConnectDB()

	router := mux.NewRouter().StrictSlash(true)

	controller := controllers.Controller{}

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/recipes", controller.GetAllRecipes(db)).Methods("GET")
	router.HandleFunc("/cuisines", controller.GetAllCuisines(db)).Methods("GET")
	router.HandleFunc("/recipe/{recipe_id}/ingredients", controller.GetRecipeIngredients(db)).Methods("GET")

	router.HandleFunc("/savedrecipes/{user_id}", controller.GetUserRecipes(db)).Methods("GET")
	router.HandleFunc("/savedrecipes/{user_id}/{recipe_id}", controller.AddUserRecipes(db)).Methods("PUT")
	router.HandleFunc("/savedrecipes/{user_id}/{recipe_id}", controller.RemoveUserRecipes(db)).Methods("DELETE")

	router.HandleFunc("/protectedEndpoint", utils.TokenVerifyMiddleWare(controller.Protected(db))).Methods("GET")
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}