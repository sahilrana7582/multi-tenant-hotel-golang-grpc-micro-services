GO := go
SERVICE_NAME := user-service
MIGRATE := migrate
MIGRATION_DIR := db/migrations
DB_URL := postgres://postgres:1234@localhost:5432/multi-tenent-hotel?sslmode=disable



.PHONY: run 

run:
	@echo "🚀 Starting Tenant Service..."
	cd cmd && go run main.go
