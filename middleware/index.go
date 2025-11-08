package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

type contextKey string

const userIDKey contextKey = "userID"

type Middleware func(http.Handler) http.Handler

type WrappedWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrappedWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Logging middleware
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &WrappedWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Printf("%s %s %d %s", r.Method, r.URL.Path, wrapped.StatusCode, time.Since(start))
	})
}

// EnsureAuthenticated middleware - validates auth header and adds userID to context
func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Simple validation - in production use proper JWT validation
		if authHeader != "Bearer valid-token" {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		// Extract userID (in production, decode from JWT)
		userID := "user-123"

		// Add userID to context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// CreateStack chains multiple middleware together
func CreateStack(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}