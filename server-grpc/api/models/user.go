package models

// User ...
type User struct {
	ID        int64  `json:"user_id"`
	FullName  string `json:"name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedOn string `json:"created_on"`
}
