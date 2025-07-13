package routes

import (
	"gateway/proxy"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	routes := map[string]string{
		"/api/": "http://localhost:8000",
	}

	for prefix, host := range routes {
		mux.Handle(prefix, http.StripPrefix(prefix, proxy.New(host)))
	}
}
