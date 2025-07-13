package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/handler"
)

func NewRouter(h *handler.TenantHandler) http.Handler {
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

	r.Route("/tenants", func(r chi.Router) {
		r.Post("/", CustomHandler(h.CreateTenant))
		r.Get("/", CustomHandler(h.ListTenants))
		r.Get("/{id}", CustomHandler(h.GetTenant))
		r.Put("/{id}", CustomHandler(h.UpdateTenant))
		r.Delete("/{id}", CustomHandler(h.DeleteTenant))
	})

	return r
}
