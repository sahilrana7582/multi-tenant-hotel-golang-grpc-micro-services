package routes

import (
	"gateway/middleware"
	"gateway/proxy"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	routes := map[string]string{
		"/api/tenants/":    "http://localhost:8000",
		"/api/users/":      "http://localhost:8001",
		"/api/department/": "http://localhost:8002",
		"/api/auth/":       "http://localhost:8004",
		"/api/role/":       "http://localhost:8005",
		"/api/permission/": "http://localhost:8005",
		"/api/rooms/":      "http://localhost:8006",
	}

	for prefix, host := range routes {

		if prefix == "/api/auth/" || prefix == "/api/tenants/" {
			mux.Handle(prefix, http.StripPrefix(prefix, proxy.New(host)))

			continue
		}

		mux.Handle(prefix, middleware.AuthMiddleware(http.StripPrefix(prefix, proxy.New(host))))
	}
}
