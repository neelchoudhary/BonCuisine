package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/neelchoudhary/boncuisine/api/controllers/recipe/recipepb"

	"github.com/neelchoudhary/boncuisine/api/driver"

	recipecontrollers "github.com/neelchoudhary/boncuisine/api/controllers/recipe"

	_ "github.com/lib/pq"

	_ "github.com/neelchoudhary/boncuisine/docs"

	"google.golang.org/grpc"
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
	fmt.Println("Starting Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//var basePath = "/api/v1"
	var env = flag.String("env", "local", "environment type: local, develop, staging, production")
	flag.Parse()

	s := grpc.NewServer()
	controller := recipecontrollers.SetDataSource(driver.ConnectDB(*env))
	recipepb.RegisterRecipeServiceServer(s, controller)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
