package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/service"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
)

func NewUserRoutes(db *pgxpool.Pool) http.Handler {
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

	hotelRepo := repo.NewHotelRepo(db)
	hotelService := service.NewHotelRepo(&hotelRepo)
	hotetlHanlder := handler.NewHotelHandler(hotelService)

	r.Route("/", func(r chi.Router) {
		r.Post("/create", responsewriter.CustomHandler(hotetlHanlder.CreateHotelInfo))
		// r.Get("/{id}", responsewriter.CustomHandler(roomHandler.GetRoomByID))
		// r.Get("/", responsewriter.CustomHandler(roomHandler.GetAllRooms))
		// r.Put("/{id}", responsewriter.CustomHandler(h.UpdateRole))
		// r.Delete("/{id}", responsewriter.CustomHandler(h.DeleteRole))

		//Location
		r.Route("/location", func(r chi.Router) {
			r.Post("/create", responsewriter.CustomHandler(hotetlHanlder.CreatNewLocation))
			r.Get("/", responsewriter.CustomHandler(hotetlHanlder.GetHotelLocation))
		})

	})

	return r
}
