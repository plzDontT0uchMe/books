package middleware

import (
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"backend/go/books/pkg/berror"
)

// RateLimiter прослойка, ограничивающая количество запросов к роутам.
func RateLimiter(maxRequests int, duration time.Duration) func(http.Handler) http.Handler {
	var mu sync.Mutex
	tokens := make(map[string]int)
	ticker := time.NewTicker(duration)

	go func() {
		for range ticker.C {
			mu.Lock()
			for ip := range tokens {
				tokens[ip] = maxRequests
			}
			mu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			ip, _, _ := net.SplitHostPort(r.RemoteAddr)

			mu.Lock()
			if _, exists := tokens[ip]; !exists {
				tokens[ip] = maxRequests
			}

			if tokens[ip] > 0 {
				tokens[ip]--
				mu.Unlock()
				next.ServeHTTP(w, r)
			} else {
				mu.Unlock()
				errNew := berror.TooManyRequests()
				slog.Log(ctx, errNew.Level(), errNew.Error())
				berror.HTTPError(w, errNew)
			}
		})
	}
}
