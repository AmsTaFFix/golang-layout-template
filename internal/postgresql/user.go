package postgresql

import (
	"context"
	"github.com/AmsTaFFix/golang-layout-template/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	pgPool *pgxpool.Pool
}

func NewUserStorage(pgPool *pgxpool.Pool) *UserStorage {
	return &UserStorage{pgPool: pgPool}
}

func (u *UserStorage) Get(ctx context.Context, userID internal.UserID) (internal.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserStorage) Set(ctx context.Context, user internal.User) error {
	//TODO implement me
	panic("implement me")
}
