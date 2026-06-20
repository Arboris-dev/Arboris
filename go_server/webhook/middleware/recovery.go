package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				slog.Error("Panic Recovered", "ERROR", err, "Stack", string(debug.Stack()))
				http.Error(w, "Unable to receive data from github", http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}
