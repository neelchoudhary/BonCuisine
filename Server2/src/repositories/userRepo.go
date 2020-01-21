package repositories

import (
	"database/sql"

	"github.com/neelchoudhary/boncuisine/models"
)

type UserRepository struct{}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where user_name=$1", user.UserName)
	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Email, &user.Password, &user.CreatedOn)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	stmt := "insert into users (user_id, name, user_name, email, password, created_on) values($1, $2, $3, $4, $5, $6) RETURNING user_id;"
	err := db.QueryRow(stmt, user.ID, user.FullName, user.UserName, user.Email, user.Password, user.CreatedOn).Scan(&user.ID)

	logFatal(err)

	user.Password = ""
	return user
}
