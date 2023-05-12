package http

import (
	"context"
	"github.com/AmsTaFFix/golang-layout-template/internal"
)

type UserRepository interface {
	Get(ctx context.Context, userId internal.UserId) (internal.User, error)
	Create(ctx context.Context, userId internal.UserId, name string, managerId internal.ManagerId) (internal.User, error)
}

type ProductRepository interface {
	Get(ctx context.Context, productId internal.ProductId) (internal.Product, error)
	Create(ctx context.Context, productId internal.ProductId, name string, managerId internal.ManagerId) (internal.Product, error)
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
