package inmemory

import (
	"context"
	"github.com/AmsTaFFix/golang-layout-template/internal"
)

type UserStorage struct {
	data map[internal.UserID]internal.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{data: map[internal.UserID]internal.User{}}
}

func (u *UserStorage) Get(ctx context.Context, userID internal.UserID) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStorage) Set(ctx context.Context, user internal.User) error {
	//TODO implement me
	panic("implement me")
}
