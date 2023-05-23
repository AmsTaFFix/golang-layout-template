package internal

//go:generate mockgen -source=user.go -destination=./mock/user.go -package=mock

import (
	"context"
	"fmt"
)

type UserId string

type User struct {
	Id        UserId
	FirstName string
	LastName  string
}

type UserStorage interface {
	Get(ctx context.Context, userId UserId) (User, error)
	Set(ctx context.Context, user User) error
}

type UserRepository struct {
	storage UserStorage
}

func NewUserRepository(storage UserStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (ur *UserRepository) Get(ctx context.Context, userId UserId) (User, error) {
	// add some rules, additional business logic, etc.
	user, err := ur.storage.Get(ctx, userId)
	if err != nil {
		return User{}, fmt.Errorf("can't get user from storage with id '%s': %w", userId, err)
	}

	return user, nil
}

func (ur *UserRepository) Create(ctx context.Context, userId UserId, name string, managerId ManagerId) (User, error) {
	// add some rules, additional business logic, etc.
	_ = managerId // add some audit?

	user := User{Id: userId, FirstName: name}
	err := ur.storage.Set(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("can't create user in storage")
	}

	return user, nil

}
