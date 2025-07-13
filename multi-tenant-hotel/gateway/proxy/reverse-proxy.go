package proxy

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func New(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Error From Gateway | New Proxy: \n", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ModifyResponse = func(res *http.Response) error {
		// Optional: you could manipulate response headers or body here
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("[GATEWAY] Reverse proxy error for %s %s: %v", r.Method, r.URL.Path, err)

		// Standard error response format
		errorResponse := map[string]interface{}{
			"status":    http.StatusServiceUnavailable,
			"error":     "Service Unavailable",
			"message":   "The service you're trying to reach is currently unavailable. Please try again later.",
			"path":      r.URL.Path,
			"method":    r.Method,
			"timestamp": time.Now().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(errorResponse)
	}

	return proxy
}
