package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/handler"
)

func NewRouter(h *handler.TenantHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/tenants", func(r chi.Router) {
		r.Post("/", CustomHandler(h.CreateTenant))
		// r.Get("/",     CustomHandler(h.ListTenants))
		// r.Get("/{id}", CustomHandler(h.GetTenant))
		// r.Put("/{id}", CustomHandler(h.UpdateTenant))
		// r.Delete("/{id}", CustomHandler(h.DeleteTenant))
	})

	return r
}
