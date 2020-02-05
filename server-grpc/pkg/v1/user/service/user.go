package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/neelchoudhary/boncuisine/db/models"
	repository "github.com/neelchoudhary/boncuisine/db/repositories"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("All fields are required"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpUser.GetPassword()), 10)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to hash password: %s", err.Error()))
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

		tokenString, err := utils.CreateToken(userToLogIn.ID)
		if err != nil {
			return nil, err
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
