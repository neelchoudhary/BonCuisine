package models

//Document schema structure
type Document struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//ToDoList example structure
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

//UserModel user structure
type UserModel struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string             `json:"username,omitempty"`
	Email      string             `json:"email,omitempty"`
	University string             `json:"university,omitempty"`
	FullName   string             `json:"fullName,omitempty"`
	Password   string             `json:"password,omitempty"`
}
