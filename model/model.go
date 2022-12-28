package model

type User struct {
	Name     string `json:"name,omitempty" bson:"name"`
	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
}
