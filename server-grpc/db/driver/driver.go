package driver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/neelchoudhary/boncuisine/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	// Required imports
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DBConnection ...
type DBConnection struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func getDBCredentials(env string) DBConnection {
	var secretName string
	if env == "local" {
		return localCredentials()
	} else if env == "develop" {
		secretName = "RDSDevelopCredentials"
	} else if env == "staging" {
		secretName = "RDSStagingCredentials"
	} else if env == "production" {
		secretName = "RDSProductionCredentials"
	} else {
		return localCredentials()
	}
	dbConnString, err := utils.GetAWSSecret(secretName)

	utils.LogIfFatalAndExit(err)

	var dbConnection DBConnection
	if dbConnString != "" {
		json.Unmarshal([]byte(dbConnString), &dbConnection)
	}
	return dbConnection
}

func localCredentials() DBConnection {
	var dbConnection DBConnection
	dbConnection.Host = "localhost"
	dbConnection.Port = 5432
	dbConnection.User = "postgres"
	dbConnection.Password = "password"
	dbConnection.DbName = "postgres"

	return dbConnection
}

// ConnectDB connect to postgresql db
func ConnectDB(env string) *sql.DB {
	fmt.Println("Connecting... " + env)
	dbConnection := getDBCredentials(env)
	fmt.Println("Retreived Credentials: " + dbConnection.Host)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConnection.Host, dbConnection.Port, dbConnection.User, dbConnection.Password, dbConnection.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	utils.LogIfFatalAndExit(err)

	err = db.Ping()
	utils.LogIfFatalAndExit(err)

	fmt.Println("You connected to your database.")

	MigrateUp(dbConnection)
	if env == "local" || env == "develop" {
		LoadSampleData(db)
	}

	return db
}

// LoadSampleData loads sample data
func LoadSampleData(db *sql.DB) {
	fmt.Println("Loading Sample Data... ")
	query, err := ioutil.ReadFile("db/sample_data/sample_data.sql")
	utils.LogIfFatalAndExit(err)

	_, err = db.Exec(string(query))
	utils.LogIfFatalAndExit(err)
	fmt.Println("Loaded Sample Data")
}

// MigrateUp migrate up to most recent migration
// Used for development and continuous integration
// migrate create -ext sql -dir db/migrations -seq create_users_table
// migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations up
// migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations down
func MigrateUp(dbConnection DBConnection) {
	fmt.Println("Migrating Up...")
	sourceURL := "file://db/migrations"
	// postgres://user:password@host:port/dbname?query
	localDbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbConnection.User, dbConnection.Password, dbConnection.Host, dbConnection.DbName)
	m, err := migrate.New(sourceURL, localDbURL)
	utils.LogIfFatalAndExit(err)

	if err := m.Up(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Finished Migrations. Up to Date.")
}
