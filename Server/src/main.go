package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"flag"

	"github.com/neelchoudhary/boncuisine/api/utils"

	"github.com/neelchoudhary/boncuisine/api/driver"

	"github.com/neelchoudhary/boncuisine/api/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/swaggo/http-swagger"
	_ "github.com/neelchoudhary/boncuisine/docs"
)

var db *sql.DB
// @title BonCuisine API
// @version 1.0.0
// @description API for the BonCuisine app.

// @contact.name BonCuisinie API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1

func main() {
	var basePath = "/api/v1"
	var env = flag.String("env", "local", "environment type, local, develop, staging, production")
	flag.Parse()
	db = driver.ConnectDB(*env)

	router := mux.NewRouter().StrictSlash(true)

	controller := controllers.Controller{}

	router.HandleFunc("/", homeLink)
	router.HandleFunc(basePath+"/recipes", controller.GetAllRecipes(db)).Methods("GET")
	router.HandleFunc(basePath+"/cuisines", controller.GetAllCuisines(db)).Methods("GET")
	router.HandleFunc(basePath+"/recipe/{recipe_id}/ingredients", controller.GetRecipeIngredients(db)).Methods("GET")

	router.HandleFunc(basePath+"/savedrecipes/{user_id}", controller.GetUserRecipes(db)).Methods("GET")
	router.HandleFunc(basePath+"/savedrecipes/{user_id}/{recipe_id}", controller.AddUserRecipes(db)).Methods("POST")
	router.HandleFunc(basePath+"/savedrecipes/{user_id}/{recipe_id}", controller.RemoveUserRecipes(db)).Methods("DELETE")

	router.HandleFunc(basePath+"/protectedEndpoint", utils.TokenVerifyMiddleWare(controller.Protected(db))).Methods("GET")
	router.HandleFunc(basePath+"/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc(basePath+"/login", controller.Login(db)).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// TODO move to controller
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}