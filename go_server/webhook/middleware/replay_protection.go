package middleware

import (
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func PreventReplay(cache *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			deliveryId := r.Header.Get("X-Github-Delivery")
			ctx := r.Context()

			ok, err := cache.SetNX(ctx, "deliver:"+deliveryId, true, time.Second*600).Result()

			if err != nil || !ok {
				http.Error(w, "The data already exists", http.StatusConflict)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
