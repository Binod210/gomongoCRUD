package services

import (
	"context"

	"github.com/Binod210/gomongoCRUD/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	u.userCollection.InsertOne(u.ctx, user)
	return nil, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*model.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) DeleteUser(userId string) error {
	return nil
}
