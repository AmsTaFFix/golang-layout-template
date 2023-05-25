package http

import (
	"context"
	"github.com/AmsTaFFix/golang-layout-template/internal"
)

type UserRepository interface {
	Get(ctx context.Context, userID internal.UserID) (internal.User, error)
	Create(ctx context.Context, userID internal.UserID, name string, managerID internal.ManagerID) (internal.User, error)
}

type ProductRepository interface {
	Get(ctx context.Context, productID internal.ProductID) (internal.Product, error)
	Create(ctx context.Context, productID internal.ProductID, name string, managerID internal.ManagerID) (internal.Product, error)
}

type Server struct {
	userRepository    UserRepository
	productRepository ProductRepository
}

func NewServer(userRepository UserRepository, productRepository ProductRepository) *Server {
	return &Server{userRepository: userRepository, productRepository: productRepository}
}

func (s *Server) Start() error {
	// ...
	return nil
}
