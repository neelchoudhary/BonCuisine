package repository

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/api/models"
	"github.com/neelchoudhary/boncuisine/pkg/utils"
)

// AccountRepository struct
type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository sets the data source (e.g database)
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Login logs into account
func (r *AccountRepository) Login(user models.User) (models.User, error) {
	row := r.db.QueryRow("select * from users where user_name=$1", user.UserName)
	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Password, &user.CreatedOn)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Signup create new account
func (r *AccountRepository) Signup(user models.User) models.User {
	stmt := "insert into users (user_id, name, user_name, email, password, created_on) values($1, $2, $3, $4, $5, $6) RETURNING user_id;"
	err := r.db.QueryRow(stmt, user.ID, user.FullName, user.UserName, user.Email, user.Password, user.CreatedOn).Scan(&user.ID)

	utils.LogFatal(err)

	user.Password = ""
	return user
}
