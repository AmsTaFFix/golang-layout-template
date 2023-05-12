package postgresql

import (
	"context"
	"github.com/AmsTaFFix/golang-layout-template/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStorage struct {
	pgPool *pgxpool.Pool
}

func NewProductStorage(pgPool *pgxpool.Pool) *ProductStorage {
	return &ProductStorage{pgPool: pgPool}
}

func (p *ProductStorage) Get(ctx context.Context, productId internal.ProductId) (internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductStorage) Set(ctx context.Context, product internal.Product) error {
	//TODO implement me
	panic("implement me")
}
