package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/neelchoudhary/boncuisine/db/models"
	repository "github.com/neelchoudhary/boncuisine/db/repositories"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type userServiceServer struct {
	userRepo *repository.UserRepository
}

// NewUserServiceServer contructor to assign repo
func NewUserServiceServer(db *sql.DB) user.UserServiceServer {
	return &userServiceServer{userRepo: repository.NewUserRepository(db)}
}

func (s *userServiceServer) Signup(ctx context.Context, req *user.SignupRequest) (*user.SignupResponse, error) {
	// Check if email already exists in db
	signUpUser := req.GetSignUpUser()
	if signUpUser.GetFullname() == "" || signUpUser.GetUsername() == "" || signUpUser.GetEmail() == "" || signUpUser.GetPassword() == "" {
		log.Fatal("All fields required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpUser.GetPassword()), 10)

	if err != nil {
		log.Fatal("Failed to hash password. ", err)
	}
	uniqueID := uuid.NewV4()
	newUser := user.User{
		Id:        uniqueID.String(),
		Email:     signUpUser.Email,
		Password:  string(hash),
		Fullname:  signUpUser.Fullname,
		Username:  signUpUser.Username,
		CreatedOn: time.Now().Format("2006-01-02T15:04:05"),
	}
	err = s.userRepo.CreateUser(*signUpPbToData(newUser))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error creating user: %s", err.Error()))
	}

	res := &user.SignupResponse{
		Success: true,
	}

	return res, nil
}

func (s *userServiceServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	loginUser := req.GetLoginUser()
	email := loginUser.GetEmail()
	password := loginUser.GetPassword()
	userToLogIn, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting user by email: %s", err.Error()))
	}

	if userToLogIn.ID != "" {
		err := bcrypt.CompareHashAndPassword([]byte(userToLogIn.Password), []byte(password))
		if err != nil {
			// Error, incorrect password
			return nil, status.Errorf(codes.PermissionDenied, fmt.Sprintf("Invalid Login Credientials: %s", err.Error()))
		}
		// Token expires in 50 minutes
		expirationTime := time.Now().Add(50 * time.Minute)
		claims := &Claims{
			UserID: userToLogIn.ID,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Login and get the encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte("verySecretSecret"))
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to sign token: %s", err.Error()))
		}
		res := &user.LoginResponse{
			Success: true,
			Token:   tokenString,
		}
		return res, nil
	}
	// User with the given email does not exist
	return nil, status.Errorf(codes.PermissionDenied, fmt.Sprintf("Invalid Login Credientials"))
}

func signUpPbToData(data user.User) *models.User {
	return &models.User{
		ID:        data.Id,
		Email:     data.Email,
		Password:  data.Password,
		FullName:  data.Fullname,
		CreatedOn: data.CreatedOn,
		UserName:  data.Username,
	}
}
