package middleware

import (
	"context"
	"net/http"

	"github.com/coding-monk-2000/auth-api/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ContextUsernameKey = contextKey("username")

// AuthMiddleware validates JWT and stores username claim in request context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := utils.ExtractTokenFromHeader(r.Header.Get("Authorization"))
		token, err := utils.ValidateToken(tokenStr)
		if err != nil || token == nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
