package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/merteldem1r/TaskeFlow-API/internal/utils"
)

type contextKey string

const (
	ContextUserID contextKey = "user_id"
	ContextRole   contextKey = "role"
)

// JWTAuth checks for a valid JWT and attaches user info to the context.
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.JSON(w, http.StatusUnauthorized, utils.APIResponse{
				Status: "error",
				Error:  "Unauthorized: missing or invalid token",
			})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			utils.JSON(w, http.StatusUnauthorized, utils.APIResponse{
				Status: "error",
				Error:  "Unauthorized: invalid token",
			})
			return
		}

		// Attach user info to context
		ctx := context.WithValue(r.Context(), ContextUserID, claims["user_id"])
		ctx = context.WithValue(ctx, ContextRole, claims["role"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
