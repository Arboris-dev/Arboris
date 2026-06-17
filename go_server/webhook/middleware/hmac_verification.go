package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
)

func VerifyHMAC(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, "Unable to read the webhook data", http.StatusInternalServerError)
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(body))
			sig := r.Header.Get("X-Hub-Signature-256")
			mac := hmac.New(sha256.New, []byte(secret))
			mac.Write(body)

			expected := "sha256=" + hex.EncodeToString(mac.Sum(nil))

			if !hmac.Equal([]byte(sig), []byte(expected)) {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
