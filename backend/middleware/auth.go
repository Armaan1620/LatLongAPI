package middleware

import (
	"context"
	"latlongapi/backend/auth"
	"latlongapi/backend/handlers"
	"latlongapi/backend/models"
	"latlongapi/backend/store"
	"net/http"
)

// AuthMiddleware validates JWT tokens and sets user in context
func AuthMiddleware(userStore models.UserStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := handlers.GetTokenFromRequest(r)
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := auth.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			user, err := userStore.GetUserByID(claims.UserID)
			if err != nil {
				if err == store.ErrUserNotFound {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Add user to context
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// OptionalAuthMiddleware validates JWT tokens if present but doesn't require them
func OptionalAuthMiddleware(userStore models.UserStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := handlers.GetTokenFromRequest(r)
			if tokenString != "" {
				claims, err := auth.ValidateToken(tokenString)
				if err == nil {
					user, err := userStore.GetUserByID(claims.UserID)
					if err == nil {
						ctx := context.WithValue(r.Context(), "user", user)
						r = r.WithContext(ctx)
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

