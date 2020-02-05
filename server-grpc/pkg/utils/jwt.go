package utils

import (
	"context"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	tokenExpiryMinutes = 50
	secretJWTSecret    = "/jwtsecret"
)

type tokenAuth struct {
	tokenString string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.tokenString,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}

// GetTokenAuth gets auth token to pass into each rpc
func GetTokenAuth(data string) tokenAuth {
	return tokenAuth{tokenString: data}
}

type claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func getJWTSecret() (string, error) {
	env, err := ReadEnvFromFile()
	if err != nil {
		return "", status.Errorf(codes.Internal, "Failed to read env from file: "+err.Error())
	}
	jwtSecret, err := GetAWSSecret(env + secretJWTSecret)
	if err != nil {
		return "", err
	}
	return jwtSecret, nil
}

// CreateToken crates a new token with a claim, returns grpc errors
func CreateToken(userID string) (string, error) {
	// Token expires in 50 minutes
	expirationTime := time.Now().Add(tokenExpiryMinutes * time.Minute)
	claims := &claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret, err := getJWTSecret()
	if err != nil {
		return "", err
	}

	// Login and get the encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", status.Errorf(codes.Internal, "Failed to sign token: "+err.Error())

	}
	return tokenString, nil
}

// AuthorizeToken authorizes the token received from metadata, returns grpc errors
func AuthorizeToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]

	const prefix = "Bearer "
	if !strings.HasPrefix(token, prefix) {
		return "", status.Error(codes.Unauthenticated, `missing "Bearer " prefix in "Authorization" header`)
	}

	token = strings.TrimPrefix(token, prefix)

	jwtSecret, err := getJWTSecret()
	if err != nil {
		return "", err
	}

	// validateToken function validates the token
	userID, err := validateToken(token, jwtSecret)

	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, err.Error())
	}
	return userID, nil
}

// ValidateToken validates the provided token string
func validateToken(tokenString string, jwtSecret string) (string, error) {
	claims := &claims{}
	// Parse the JWT string and store the result in claims.
	// This method will return an error if the token is invalid/expired,
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
