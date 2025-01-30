package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"backend/go/books/pkg/berror"
)

// Recover прослойка, которая перехватывает паники, возникающие в обработчиках HTTP-запросов,
// и возвращает клиенту внутреннюю ошибку сервера. Она также логирует информацию о панике,
// включая стек вызовов, с использованием slog.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				berror.HTTPError(w, berror.Internal().Obj("recover"))
				slog.ErrorContext(r.Context(), "panic", "error", rec, "stack_trace", string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
