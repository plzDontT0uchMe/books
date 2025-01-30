package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Log прослойка, которая логирует детали каждого HTTP-запроса и его ответа.
// Она логирует метод, путь URL, код состояния и продолжительность запроса.
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		start := time.Now()
		slog.DebugContext(ctx, "incoming request", "method", r.Method, "req", r.RequestURI)

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		defer func() {
			duration := time.Since(start)
			slog.InfoContext(ctx, "completed request", "remote_address", r.RemoteAddr, "method", r.Method,
				"req", r.RequestURI, "status", rw.statusCode, "duration", duration)
		}()

		next.ServeHTTP(rw, r)
	})
}
