package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func ExtractTokenFromHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}
	a := strings.TrimSpace(authHeader)
	if strings.HasPrefix(strings.ToLower(a), "bearer ") {
		return strings.TrimSpace(a[7:])
	}
	return a
}

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses HMAC (HS256/HS384/HS512)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return jwtKey, nil
		}
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})
}
