package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		
		slog.Info("Webhook Processed",
			"method", r.Method,
			"path", r.URL.Path,
			"deliveryId", r.Header.Get("X-Github-Delivery"),
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
