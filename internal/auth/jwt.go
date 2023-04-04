package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJWT(id uuid.UUID) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": id,
	}).SignedString([]byte("secret"))
}

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, ok := r.Header["Authorization"]
		if len(t) == 0 {
			log.Println("Token not found")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(t[0], " ")[1]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			log.Println(err.Error(), token)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("Failed to validate")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", claims["uid"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
