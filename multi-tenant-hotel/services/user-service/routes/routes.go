package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/handler"
)

func NewRouter(h *handler.UserHandler) http.Handler {
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

	r.Route("/", func(r chi.Router) {
		r.Post("/create", responsewriter.CustomHandler(h.CreateUser))
		// r.Get("/{id}", CustomHandler(h.GetUser))
		// r.Put("/{id}", CustomHandler(h.UpdateUser))
		// r.Delete("/{id}", CustomHandler(h.DeleteUser))
	})

	return r
}
