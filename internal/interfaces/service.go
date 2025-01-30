package interfaces

import (
	"context"

	"backend/go/books/internal/dto"
)

type Service interface {
	GetBooks(ctx context.Context) (dto.Books, error)
	GetAuthors(ctx context.Context) (dto.Authors, error)
	GetBook(ctx context.Context, id int) (*dto.Book, error)
}
