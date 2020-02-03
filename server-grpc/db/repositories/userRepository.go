package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/db/models"
)

// UserRepository struct
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository sets the data source (e.g database)
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Login tries to find user with matching credentials
func (r *UserRepository) Login(user models.User) (models.User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE user_name=$1", user.UserName)
	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Password, &user.CreatedOn)

	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByEmail tries to find user with email
func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row := r.db.QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Password, &user.CreatedOn)
	if err == sql.ErrNoRows {
		return user, nil
	} else if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser create new user for signup
func (r *UserRepository) CreateUser(user models.User) error {
	var userID string
	statement := "INSERT INTO users (user_id, name, user_name, email, password, created_on) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id;"
	err := r.db.QueryRow(statement, user.ID, user.FullName, user.UserName, user.Email, user.Password, user.CreatedOn).Scan(&userID)

	if err != nil {
		return err
	}

	return nil
}
