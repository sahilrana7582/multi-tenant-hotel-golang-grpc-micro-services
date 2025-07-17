package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/role-service/config"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/db"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/routes"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/service"
)

func main() {

	cfg := config.LoadConfig("../.env")

	db, err := db.NewPostgresPool(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repo.NewRoleRepo(db)
	service := service.NewRoleService(repo)
	hand := handler.NewRoleHandler(service)
	r := routes.NewRouter(hand)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.SERVER_PORT),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 Role service started on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
