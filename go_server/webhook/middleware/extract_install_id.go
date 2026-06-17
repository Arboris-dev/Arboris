package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type contextKey string

const installationIDKey contextKey = "installationID"

func ExtractInstallID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewReader(body))

		var payload struct {
			Installation struct {
				ID int64 `json:"id"`
			} `json:"installation"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), installationIDKey, strconv.FormatInt(payload.Installation.ID, 10))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
