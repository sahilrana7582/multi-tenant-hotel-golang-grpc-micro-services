package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/room-service/config"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/db"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/routes"
)

func main() {
	cfg := config.LoadConfig("../.env")

	db, err := db.NewPostgresPool(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := routes.NewRouter(db)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.SERVER_PORT),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 Room service started on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
