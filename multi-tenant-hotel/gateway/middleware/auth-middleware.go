package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sahilrana7582/multi-tenant-hotel/pkg/auth"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			responsewriter.WriteError(w, errs.New("ERR_UNAUTHORIZED", "Authorization header missing", http.StatusUnauthorized))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ParseJWT(tokenStr)
		if err != nil {
			responsewriter.WriteError(w, errs.New("ERR_INVALID_TOKEN", fmt.Sprintf("Reason: %s", err.Error()), http.StatusUnauthorized))
			return
		}

		tenantID := claims.TenantID
		userID := claims.UserID

		r.Header.Set("X-Tenant-ID", tenantID)
		r.Header.Set("X-User-ID", userID)

		next.ServeHTTP(w, r)
	})
}
