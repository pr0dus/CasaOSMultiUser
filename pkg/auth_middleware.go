package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/IceWhaleTech/CasaOS/model"
)

var jwtKey = []byte("your_secret_key") // Replace with your actual secret key

// Claims struct to store user claims
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// AuthMiddleware validates the JWT token and retrieves user info from it
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}