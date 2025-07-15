package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/config"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/db"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/routes"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/service"
)

func main() {
	cfg := config.LoadConfig("../.env")

	db, err := db.NewPostgresPool(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	authRepo := repo.NewAuthRepo(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := routes.NewRouter(authHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.SERVER_PORT),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 Auth service started on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}

}
