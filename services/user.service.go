package services

import "github.com/Binod210/gomongoCRUD/model"

type UserService interface {
	CreateUser(*model.User) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(string) error
}
