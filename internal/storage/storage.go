package storage

import (
	"context"

	"backend/go/books/internal/models"
	"backend/go/books/pkg/berror"
)

// Слой базы данных (в данном случае - хранилище в памяти)

type Storage struct {
}

// New создает новое хранилище.
func New() *Storage {
	return &Storage{}
}

// GetBooks возвращает список книг.
func (s *Storage) GetBooks(ctx context.Context) (models.Books, error) {
	return Books, nil
}

// GetAuthors возвращает список авторов.
func (s *Storage) GetAuthors(ctx context.Context) (models.Authors, error) {
	return Authors, nil
}

// GetBook возвращает книгу по идентификатору.
func (s *Storage) GetBook(ctx context.Context, id int) (*models.Book, error) {
	for _, b := range Books {
		if b.Id == id {
			return &b, nil
		}
	}
	return nil, berror.NotFound().Obj("book")
}
