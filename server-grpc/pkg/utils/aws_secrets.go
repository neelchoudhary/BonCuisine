package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAWSSecret retrieves secret from AWS Secret Manager, returns grpc errors
func GetAWSSecret(secretName string) (string, error) {
	region := "us-east-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", status.Errorf(codes.Internal, "Failed to get AWS Secret: "+err.Error())
	}

	if result.SecretString == nil {
		return "", status.Errorf(codes.Internal, "Failed to get AWS Secret: Empty secret")
	}
	return *result.SecretString, nil
}
