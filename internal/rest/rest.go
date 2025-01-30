package rest

import (
	"fmt"
	"net/http"

	"backend/go/books/internal/config"
	"backend/go/books/internal/interfaces"
	"backend/go/books/internal/rest/author"
	"backend/go/books/internal/rest/book"
	"backend/go/books/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

// New создает http сервер с routes и middleware.
func New(cfg config.Server, sv interfaces.Service) http.Server {
	r := chi.NewRouter()

	r.Use(middleware.CORS)
	r.Use(middleware.Recover)
	r.Use(middleware.ReqId)
	r.Use(middleware.Log)
	r.Use(middleware.RateLimiter(cfg.RateLimitCount, cfg.RateLimitDuration))
	r.Use(middleware.Timeout(cfg.RequestTimeout))

	r.Get("/liveness", func(http.ResponseWriter, *http.Request) {})
	r.Get("/books", book.List(sv))
	r.Get("/authors", author.List(sv))
	r.Get("/book/{id}", book.Get(sv))

	h := http.Handler(r)

	return http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: h,
	}
}
