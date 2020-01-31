package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/neelchoudhary/boncuisine/api/models"
	user "github.com/neelchoudhary/boncuisine/pkg/v1/user/api"
	"github.com/neelchoudhary/boncuisine/pkg/v1/user/repository"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type userServiceServer struct {
	userRecipeRepo *repository.UserRecipeRepository
	userRepo       *repository.UserRepository
}

// NewUserServiceServer TODO
func NewUserServiceServer(db *sql.DB) user.UserServiceServer {
	return &userServiceServer{userRecipeRepo: repository.NewUserRecipeRepository(db), userRepo: repository.NewUserRepository(db)}
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
	s.userRepo.CreateUser(*signUpPbToData(newUser))

	if err != nil {
		log.Fatal(err)
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
	userToLogIn, err := s.userRepo.ContainsUser(email)
	if err != nil {
		log.Fatal(err)
	}
	if userToLogIn.ID != "" {
		err := bcrypt.CompareHashAndPassword([]byte(userToLogIn.Password), []byte(password))
		if err != nil {
			// Error, incorrect password
			log.Fatal("Incorrect password"
		} else {
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":    userToLogIn.ID,
				"email": userToLogIn.Email,
				"name":  userToLogIn.FullName,
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte("verySecretSecret"))
			if err != nil {
				log.Fatal("Failed to sign token: ", err)
			}
			res := &user.LoginResponse{
				Success: true,
				Token:   tokenString,
			}
			return res, nil
		}
	} else {
		// Error, user does not exist
	}

	res := &user.LoginResponse{
		Success: false,
	}

	return res, nil
}

func (s *userServiceServer) GetSavedRecipes(ctx context.Context, req *user.GetSavedRecipiesRequest) (*user.GetSavedRecipiesResponse, error) {
	fmt.Printf("Get saved recipes was invoked.\n")
	userID := req.GetUserId()
	recipes := s.userRecipeRepo.GetUserRecipes(userID)
	var pbSavedRecipes []*user.SavedRecipe
	for _, recipe := range recipes {
		pbSavedRecipes = append(pbSavedRecipes, dataToSavedRecipePb(recipe))
	}
	res := &user.GetSavedRecipiesResponse{
		SavedRecipes: pbSavedRecipes,
	}
	return res, nil
}

func (s *userServiceServer) AddSavedRecipe(ctx context.Context, req *user.AddSavedRecipeRequest) (*user.AddSavedRecipeResponse, error) {
	fmt.Printf("Add saved recipe was invoked.")
	userID := req.GetUserId()
	recipeID := req.GetRecipeId()
	s.userRecipeRepo.AddUserRecipe(userID, recipeID)
	res := &user.AddSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func (s *userServiceServer) RemoveSavedRecipe(ctx context.Context, req *user.RemoveSavedRecipeRequest) (*user.RemoveSavedRecipeResponse, error) {
	fmt.Printf("Remove saved recipe was invoked.")
	userID := req.GetUserId()
	recipeID := req.GetRecipeId()
	s.userRecipeRepo.RemoveUserRecipe(userID, recipeID)
	res := &user.RemoveSavedRecipeResponse{
		Success: true,
	}
	return res, nil
}

func dataToSavedRecipePb(data models.SavedRecipe) *user.SavedRecipe {
	return &user.SavedRecipe{
		UserId:   data.UserID,
		RecipeId: data.RecipeID,
	}
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
