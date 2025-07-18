package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/service"
)

func NewRouter(db *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    http.StatusNotFound,
			"error":     "Not Found",
			"message":   "The requested resource could not be found",
			"path":      r.URL.Path,
			"method":    r.Method,
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// Room

	roomRepo := repo.NewRoomRepo(db)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	r.Route("/", func(r chi.Router) {
		r.Post("/create", responsewriter.CustomHandler(roomHandler.CreateRoom))
		r.Get("/{id}", responsewriter.CustomHandler(roomHandler.GetRoomByID))
		r.Get("/", responsewriter.CustomHandler(roomHandler.GetAllRooms))
		// r.Put("/{id}", responsewriter.CustomHandler(h.UpdateRole))
		// r.Delete("/{id}", responsewriter.CustomHandler(h.DeleteRole))
	})

	return r
}
