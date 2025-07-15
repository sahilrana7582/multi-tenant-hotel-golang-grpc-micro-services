package routes

import (
	"gateway/proxy"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	routes := map[string]string{
		"/api/tenants/":    "http://localhost:8000",
		"/api/users/":      "http://localhost:8001",
		"/api/department/": "http://localhost:8002",
	}

	for prefix, host := range routes {
		mux.Handle(prefix, http.StripPrefix(prefix, proxy.New(host)))
	}
}
