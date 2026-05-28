package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tu-usuario/taskflow/core"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				core.RespondError(w, http.StatusUnauthorized, "Token requerido")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				core.RespondError(w, http.StatusUnauthorized, "Formato de token inválido")
				return
			}

			tokenString := parts[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, core.ErrUnauthorized
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				core.RespondError(w, http.StatusUnauthorized, "Token inválido o expirado")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				core.RespondError(w, http.StatusUnauthorized, "Token claims inválidos")
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				core.RespondError(w, http.StatusUnauthorized, "Token no contiene user_id")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, int(userIDFloat))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(UserIDKey).(int)
	return id, ok
}
