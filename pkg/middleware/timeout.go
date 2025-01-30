package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"backend/go/books/pkg/berror"
)

// Timeout прослойка, которая устанавливает таймаут для каждого запроса.
func Timeout(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			r = r.WithContext(ctx)

			done := make(chan struct{})
			go func() {
				next.ServeHTTP(w, r)
				close(done)
			}()

			select {
			case <-ctx.Done():
				errNew := berror.Timeout()
				slog.Log(ctx, errNew.Level(), errNew.Error())
				berror.HTTPError(w, errNew)
			case <-done:
			}
		})
	}
}
