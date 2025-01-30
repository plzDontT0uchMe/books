package book

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"backend/go/books/internal/interfaces"
	"backend/go/books/pkg/berror"
	"github.com/go-chi/chi/v5"
)

// Слой транспортный

// Get возвращает книгу по идентификатору.
func Get(sv interfaces.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		id, err := strconv.Atoi(chi.URLParam(req, "id"))
		if err != nil {
			errNew := berror.InvalidArgument().Obj("id")
			slog.Log(ctx, errNew.Level(), errNew.Error())
			berror.HTTPError(w, errNew)
			return
		}

		book, err := sv.GetBook(ctx, id)
		if err != nil {
			berror.HTTPError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if errEncode := json.NewEncoder(w).Encode(book); errEncode != nil {
			slog.ErrorContext(ctx, errEncode.Error())
			berror.HTTPError(w, errEncode)
		}
	}
}
