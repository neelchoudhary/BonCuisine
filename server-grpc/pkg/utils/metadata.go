package utils

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// GetUserIDMetadata retrieves userID from context metadata, returns grpc errors
func GetUserIDMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.DataLoss, "Failed to get userID metadata")
	}
	if userIDMap, ok := md["userid"]; ok {
		return userIDMap[0], nil
	}
	return "", status.Errorf(codes.DataLoss, "Failed to get userID metadata")
}

// PassUserIDMetadata passes the given userID into context metadata, returns new context
func PassUserIDMetadata(ctx context.Context, userID string) context.Context {
	md := metadata.Pairs("userid", userID)
	return metadata.NewIncomingContext(ctx, md)
}
