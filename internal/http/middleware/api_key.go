package middleware

import (
	"net/http"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.Header.Get("X-API-Key")

		if key == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		if key != "super-secret-key" {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
