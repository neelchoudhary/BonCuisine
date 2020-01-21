package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dormeter.com/models"

	"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "hmm"
	dbname   = "hmm"
  )


//Database Name
dbName = "dormProject"

const urlENCODED = "application/x-www-form-urlencoded"

//PostgreSql Connection String
connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
"password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)



//Create connection with mongoDB
func init() {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Successfully connected!")
}

//AddUser adds a new user to the database
func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", urlENCODED)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Acces-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&user)
	addUser(user)
	json.NewEncoder(w).Encode(user)
}

//addUser inserts a new user to the database
func addUser(user models.UserModel) {
	result, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single User!", result.InsertedID)
}

//userExists checks if a user already exists
func userExists(user models.UserModel) bool {
	result := collection.FindOne(context.Background(), user)
	if result == nil {
		return false
	}
	return true
}
