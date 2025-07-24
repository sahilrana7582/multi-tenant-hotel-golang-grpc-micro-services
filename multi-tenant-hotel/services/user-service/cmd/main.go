package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/user-service/config"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/db"
	grpcserver "github.com/sahilrana7582/multi-tenant-hotel/user-service/grpc"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/routes"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/service"
)

func main() {

	cfg := config.LoadConfig("../.env")

	dbPool, err := db.NewPostgresPool(cfg)

	if err != nil {
		panic(err)
	}

	repo := repo.NewUserRepo(dbPool)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	r := routes.NewRouter(handler)

	go func() {
		fmt.Println("Running The Grpc Service")
		grpcserver.StartGRPCServer(service, "50051")
	}()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.SERVER_PORT),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 User service started on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
