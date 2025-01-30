package book

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"backend/go/books/internal/interfaces"
	"backend/go/books/pkg/berror"
)

// Слой транспортный

// List возвращает список книг.
func List(sv interfaces.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		books, err := sv.GetBooks(ctx)
		if err != nil {
			berror.HTTPError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if errEncode := json.NewEncoder(w).Encode(books); errEncode != nil {
			slog.ErrorContext(ctx, errEncode.Error())
			berror.HTTPError(w, errEncode)
		}
	}
}
