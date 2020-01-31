package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/api/models"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
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
	row := r.db.QueryRow("select * from users where user_name=$1", user.UserName)
	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Password, &user.CreatedOn)

	if err != nil {
		return user, err
	}

	return user, nil
}

// ContainsUser tries to find user with email
func (r *UserRepository) ContainsUser(email string) (models.User, error) {
	var user models.User
	row := r.db.QueryRow("select * from users where email=$1", email)
	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Password, &user.CreatedOn)
	if err == sql.ErrNoRows {
		return user, nil
	} else if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser create new user for signup
func (r *UserRepository) CreateUser(user models.User) {
	var userID string
	stmt := "insert into users (user_id, name, user_name, email, password, created_on) values ($1, $2, $3, $4, $5, $6) RETURNING user_id;"
	err := r.db.QueryRow(stmt, user.ID, user.FullName, user.UserName, user.Email, user.Password, user.CreatedOn).Scan(&userID)

	utils.LogFatal(err)
}
