package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("hby5jynllqzqms3d")

func NewClaims(uuid string) jwt.StandardClaims {
	return jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		Id:        uuid,
	}
}

func JWTGenerate(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JWTParse(jwtStr string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		jwtStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
