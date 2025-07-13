package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		log.Printf("Reverse proxy error: %v", err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}

	return proxy
}
