package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/config"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/db"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/handler"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/repo"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/routes"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/service"
)

func main() {

	cofig := config.LoadConfig("../.env")

	pool, err := db.NewPostgresPool(cofig)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	defer pool.Close()

	repo := repo.NewTenantRepo(pool)
	svc := service.NewTenantService(repo)
	h := handler.NewTenantHandler(*svc)

	r := routes.NewRouter(h)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cofig.ServerPort),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 Tenant service started on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
