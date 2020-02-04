package utils

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// LogFatal logs a fatal error
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetUserID retrieves userID from metadata
func GetUserID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.DataLoss, "Failed to get metadata")
	}
	if userIDMap, ok := md["userid"]; ok {
		return userIDMap[0], nil
	}
	return "", status.Errorf(codes.DataLoss, "Failed to get metadata")
}
