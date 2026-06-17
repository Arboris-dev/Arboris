package middleware

import (
	"Arboris/go_server/httpwriters"
	"io"
	"log/slog"
	"net/http"
)

func MaxBodySize(maxSize int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, maxSize)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					slog.Error("Github payload exceeds max size", "ERROR", err)
					httpwriters.RespondWithErr(w, "File exceeds max size", 413)
				}
			}(r.Body)
			next.ServeHTTP(w, r)
		})
	}
}
