package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"flag"

	"github.com/neelchoudhary/boncuisine/utils"

	"github.com/neelchoudhary/boncuisine/driver"

	"github.com/neelchoudhary/boncuisine/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	var env = flag.String("env", "local", "environment type, local, develop, staging, production")
	flag.Parse()
	db = driver.ConnectDB(*env)

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
