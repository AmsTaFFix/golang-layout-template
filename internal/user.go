package internal

//go:generate mockgen -source=user.go -destination=./mock/user.go -package=mock

import (
	"context"
	"fmt"
)

type UserID string

type User struct {
	ID        UserID
	FirstName string
	LastName  string
}

type UserStorage interface {
	Get(ctx context.Context, userID UserID) (User, error)
	Set(ctx context.Context, user User) error
}

type UserRepository struct {
	storage UserStorage
}

func NewUserRepository(storage UserStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (ur *UserRepository) Get(ctx context.Context, userID UserID) (User, error) {
	// add some rules, additional business logic, etc.
	user, err := ur.storage.Get(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("can't get user from storage with id '%s': %w", userID, err)
	}

	return user, nil
}

func (ur *UserRepository) Create(ctx context.Context, userID UserID, name string, managerID ManagerID) (User, error) {
	// add some rules, additional business logic, etc.
	_ = managerID // add some audit?

	user := User{ID: userID, FirstName: name}
	err := ur.storage.Set(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("can't create user in storage")
	}

	return user, nil

}
