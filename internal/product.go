package internal

//go:generate mockgen -source=product.go -destination=./mock/product.go -package=mock

import (
	"context"
	"fmt"
)

type ProductId string
type Product struct {
	Id   ProductId
	Name string
}
type ProductStorage interface {
	Get(ctx context.Context, productId ProductId) (Product, error)
	Set(ctx context.Context, product Product) error
}

type ProductRepository struct {
	storage ProductStorage
}

func NewProductRepository(storage ProductStorage) *ProductRepository {
	return &ProductRepository{storage: storage}
}

func (pr *ProductRepository) Get(ctx context.Context, productId ProductId) (Product, error) {
	// add some rules, additional business logic, etc.
	product, err := pr.storage.Get(ctx, productId)
	if err != nil {
		return Product{}, fmt.Errorf("can't get product from storage with id '%s': %w", productId, err)
	}

	return product, nil
}

func (pr *ProductRepository) Create(ctx context.Context, productId ProductId, name string, managerId ManagerId) (Product, error) {
	// add some rules, additional business logic, etc.
	_ = managerId // add some audit?

	product := Product{Id: productId, Name: name}
	err := pr.storage.Set(ctx, product)
	if err != nil {
		return Product{}, fmt.Errorf("can't create product in storage")
	}

	return product, nil

}
