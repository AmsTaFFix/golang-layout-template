package internal

//go:generate mockgen -source=product.go -destination=./mock/product.go -package=mock

import (
	"context"
	"fmt"
)

type ProductID string
type Product struct {
	ID   ProductID
	Name string
}
type ProductStorage interface {
	Get(ctx context.Context, productID ProductID) (Product, error)
	Set(ctx context.Context, product Product) error
}

type ProductRepository struct {
	storage ProductStorage
}

func NewProductRepository(storage ProductStorage) *ProductRepository {
	return &ProductRepository{storage: storage}
}

func (pr *ProductRepository) Get(ctx context.Context, productID ProductID) (Product, error) {
	// add some rules, additional business logic, etc.
	product, err := pr.storage.Get(ctx, productID)
	if err != nil {
		return Product{}, fmt.Errorf("can't get product from storage with id '%s': %w", productID, err)
	}

	return product, nil
}

func (pr *ProductRepository) Create(ctx context.Context, productID ProductID, name string, managerID ManagerID) (Product, error) {
	// add some rules, additional business logic, etc.
	_ = managerID // add some audit?

	product := Product{ID: productID, Name: name}
	err := pr.storage.Set(ctx, product)
	if err != nil {
		return Product{}, fmt.Errorf("can't create product in storage")
	}

	return product, nil
}
