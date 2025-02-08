package middleware

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

func RateLimiter(rps int, burst int) func(http.Handler) http.Handler {
	var (
		clients = make(map[string]*rate.Limiter)
		mu      sync.Mutex
	)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr

			mu.Lock()
			limiter, exists := clients[ip]
			if !exists {
				limiter = rate.NewLimiter(rate.Limit(rps), burst)
				clients[ip] = limiter
			}
			mu.Unlock()

			if !limiter.Allow() {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
