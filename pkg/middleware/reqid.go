package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"
const RequestID = "request_id"

// ReqId прослойка, прокидывающая в контекст идентификатор запроса из заголовка. Если заголовок был пуст или
// содержал не валидный UUID, генерируется новый идентификатор запроса.
func ReqId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		requestId, err := uuid.Parse(r.Header.Get(RequestIDHeader))
		if err != nil {
			requestId = uuid.New()
		}

		r = r.WithContext(context.WithValue(r.Context(), RequestID, requestId))

		next.ServeHTTP(rw, r)

		return
	})
}
