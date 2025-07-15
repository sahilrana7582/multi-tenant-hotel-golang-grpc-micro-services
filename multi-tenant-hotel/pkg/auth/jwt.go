package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKet = []byte("My Super Secret Key | BuLu LuLu KuLu MuLu SuLu")

type CustomClaims struct {
	TenantID string `json:"tenant_id"`
	UserID   string `json:"user_id"`
	jwt.Claims
}

func GenerateJWT(userID, tenantID string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		TenantID: tenantID,
		Claims: jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKet)
}

func ParseJWT(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKet, nil
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
