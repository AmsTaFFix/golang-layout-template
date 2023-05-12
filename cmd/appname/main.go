package main

import (
	"github.com/AmsTaFFix/golang-layout-template/internal"
	"github.com/AmsTaFFix/golang-layout-template/internal/api/grpc"
	"github.com/AmsTaFFix/golang-layout-template/internal/api/http"
	"github.com/AmsTaFFix/golang-layout-template/internal/inmemory"
	"github.com/AmsTaFFix/golang-layout-template/internal/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

func main() {
	var cfg Config // get config via flags or viper
	var productStorage internal.ProductStorage
	var userStorage internal.UserStorage
	if cfg.UseInmemory {
		productStorage = inmemory.NewProductStorage()
		userStorage = inmemory.NewUserStorage()
	} else {
		// get connect to database
		var pgxPool *pgxpool.Pool
		productStorage = postgresql.NewProductStorage(pgxPool)
		userStorage = postgresql.NewUserStorage(pgxPool)
	}

	productRepo := internal.NewProductRepository(productStorage)
	userRepo := internal.NewUserRepository(userStorage)
	grpcServer := grpc.NewServer(userRepo, productRepo)
	httpServer := http.NewServer(userRepo, productRepo)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := grpcServer.Start()
		if err != nil {
			log.Printf("can't start grpc server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := httpServer.Start()
		if err != nil {
			log.Printf("can't start http server: %s", err)
		}
	}()

	wg.Wait()
}
