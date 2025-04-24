package jwtutils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id, secret string) (string, error) {
	claims := jwt.MapClaims{
		"userid": id,
		"exp":    time.Now().Add(time.Hour * 24 * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString, secret string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

func GenerateRefreshToken(id uint, secret string) (string, error) {
	claims := jwt.MapClaims{
		"userid": id,
		"exp":    time.Now().Add(time.Hour * 24 * 14).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}