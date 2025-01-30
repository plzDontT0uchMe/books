package interfaces

import (
	"context"

	"backend/go/books/internal/models"
)

// Storage интерфейс для работы с хранилищем

type Storage interface {
	GetBooks(ctx context.Context) (models.Books, error)
	GetAuthors(ctx context.Context) (models.Authors, error)
	GetBook(ctx context.Context, id int) (*models.Book, error)
}
