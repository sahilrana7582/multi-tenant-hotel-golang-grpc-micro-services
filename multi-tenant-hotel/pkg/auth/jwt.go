package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	TenantIDKey = "X-Tenant-ID"
	UserIDKey   = "X-User-ID"
)

var jwtSecretKey = []byte("My Super Secret Key | BuLu LuLu KuLu MuLu SuLu")

type CustomClaims struct {
	TenantID string `json:"tenant_id"`
	UserID   string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, tenantID string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func ParseJWT(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetTenantID(r *http.Request) string {
	return r.Header.Get(TenantIDKey)
}

func GetUserID(r *http.Request) string {
	return r.Header.Get(UserIDKey)
}
