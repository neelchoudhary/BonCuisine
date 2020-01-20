package driver

import (
	"database/sql"
	"log"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "test"
)

var db *sql.DB

func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

// Connect to postgresql db
func ConnectDB() *sql.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	logFatal(err)


	err = db.Ping()
	logFatal(err)
	

	fmt.Println("You connected to your database.")

	return db
}

