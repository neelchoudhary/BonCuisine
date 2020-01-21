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

func getRDSCredentials() DBConnection {
	secretName := "BoncuisineRDSCredentials"
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

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connect to postgresql db
func ConnectDB() *sql.DB {
	fmt.Println("Connecting... ")
	dbConnection := getRDSCredentials()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConnection.Host, dbConnection.Port, dbConnection.User, dbConnection.Password, dbName)

	db, err = sql.Open("postgres", psqlInfo)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	fmt.Println("You connected to your database.")

	return db
}

// DBConnection ...
type DBConnection struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbInstanceIdentifier"`
}

const dbName = "boncuisinepgsql"
