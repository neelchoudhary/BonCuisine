package driver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

/// AWS login
// $(aws ecr get-login --no-include-email --region us-east-2)

/// Build new docker image on go-lang-server ECR
/// docker build -t go-lang-server .

/// Tag new docker image on go-lang-server ECR as latest
// docker tag go-lang-server:latest 729017073046.dkr.ecr.us-east-2.amazonaws.com/go-lang-server:latest

/// Push new docker image to ECR
// docker push 729017073046.dkr.ecr.us-east-2.amazonaws.com/go-lang-server:latest

/// Register New Task Definition
// aws ecs register-task-definition --family boncuisine-production-definition --requires-compatibilities FARGATE --cpu 256 --memory 512 --cli-input-json file://boncuisine-task-definition-production.json --region "us-east-2"

/// Update service with new task and start task. This should end old task
// aws ecs update-service --cluster golang-cluster --service golang-container-prod-service --task-definition boncuisine-production-definition --region "us-east-2"

func getDBCredentials(env string) DBConnection {
	var secretName string
	if (env == "local") {
		return getLocalCredentials()
	} else if (env == "develop") {
		secretName = "RDSDevelopCredentials"
	} else if (env == "staging") {
		secretName = "RDSStagingCredentials"
	} else if (env == "production") {
		secretName = "RDSProductionCredentials"
	} else {
		return getLocalCredentials()
	}
	region := "us-east-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		logFatal(err.(awserr.Error))
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString string
	var dBConnection DBConnection
	if result.SecretString != nil {
		secretString = *result.SecretString
		json.Unmarshal([]byte(secretString), &dBConnection)
	}
	return dBConnection
}

func getLocalCredentials() DBConnection {
	var dbConnection DBConnection
	dbConnection.Host = "localhost"
	dbConnection.Port = 5432
	dbConnection.User = "postgres"
	dbConnection.Password = "password"
	dbConnection.DbName = "postgres"

	return dbConnection
}

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connect to postgresql db
func ConnectDB(env string) *sql.DB {
	fmt.Println("Connecting... ")
	dbConnection := getDBCredentials(env)
	fmt.Println("Retreived Credentials")
	fmt.Println(dbConnection.Host)
	fmt.Println(dbConnection.User)
	fmt.Println(dbConnection.DbName)
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConnection.Host, dbConnection.Port, dbConnection.User, dbConnection.Password, dbConnection.DbName)

	db, err = sql.Open("postgres", psqlInfo)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	fmt.Println("You connected to your database.")

	MigrateUp(dbConnection)

	return db
}

// Used for development and continuous integration
// migrate create -ext sql -dir db/migrations -seq create_users_table
// migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations up
// migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations down
func MigrateUp(dbConnection DBConnection) {
	fmt.Println("Migrating Up...")
	sourceUrl := "file://./db/migrations"
	// postgres://user:password@host:port/dbname?query
	localDbUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", 
	dbConnection.User, dbConnection.Password, dbConnection.Host, dbConnection.DbName)
	m, err := migrate.New(sourceUrl, localDbUrl)
	if err != nil {
		logFatal(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Finished Migrations. Up to Date.")
}

// DBConnection ...
type DBConnection struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}