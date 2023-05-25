package inmemory

import (
	"context"
	"github.com/AmsTaFFix/golang-layout-template/internal"
)

type ProductStorage struct {
	data map[internal.ProductID]internal.Product
}

func NewProductStorage() *ProductStorage {
	return &ProductStorage{data: map[internal.ProductID]internal.Product{}}
}

func (p *ProductStorage) Get(ctx context.Context, productID internal.ProductID) (internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductStorage) Set(ctx context.Context, product internal.Product) error {
	//TODO implement me
	panic("implement me")
}
