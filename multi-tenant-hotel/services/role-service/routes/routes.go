package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/handler"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/repo"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/service"
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

	// Role routes
	roleRepo := repo.NewRoleRepo(db)
	roleService := service.NewRoleService(roleRepo)
	h := handler.NewRoleHandler(roleService)

	r.Route("/", func(r chi.Router) {
		r.Post("/create", responsewriter.CustomHandler(h.CreateRole))
		r.Get("/{id}", responsewriter.CustomHandler(h.GetRoleByID))
		r.Get("/", responsewriter.CustomHandler(h.GetAllRoles))
		r.Put("/{id}", responsewriter.CustomHandler(h.UpdateRole))
		r.Delete("/{id}", responsewriter.CustomHandler(h.DeleteRole))
	})

	// Permission routes
	permissionRepo := repo.NewPermissionRepo(db)
	permissionService := service.NewPermissionService(permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionService)

	r.Route("/permission", func(r chi.Router) {
		r.Post("/give", responsewriter.CustomHandler(permissionHandler.GivePermissionToRole))
	})

	return r
}
