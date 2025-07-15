package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/department-service/config"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/db"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/routes"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/service"
)

func main() {
	fmt.Println("Running The Department Service")

	cfg := config.LoadConfig("../.env")

	dbPool, err := db.NewPostgresPool(cfg)
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	repository := repo.NewDepartmentRepo(dbPool)
	departmentService := service.NewDepartmentService(repository)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	r := routes.NewRouter(departmentHandler)

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
