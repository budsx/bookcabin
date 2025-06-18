package middleware

import (
	"context"
	"net/http"

	"github.com/budsx/bookcabin/util/logger"
	"github.com/google/uuid"
)

// RequestIDMiddleware adds a unique request ID to each request context
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a new UUID for the request
		requestID := uuid.New().String()

		// Add the request ID to the response header
		w.Header().Set("X-Request-ID", requestID)

		// Add the request ID to the request context
		ctx := context.WithValue(r.Context(), logger.RequestIDKey, requestID)

		// Pass the request with the new context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
