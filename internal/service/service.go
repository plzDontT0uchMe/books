package service

import (
	"context"
	"log/slog"
	"math"

	"backend/go/books/internal/dto"
	"backend/go/books/internal/interfaces"
	"backend/go/books/pkg/berror"
)

// Сервисный слой (бизнес-логика)

type Service struct {
	st interfaces.Storage
}

// New создает новый сервис.
func New(st interfaces.Storage) *Service {
	return &Service{st: st}
}

// GetBooks возвращает список книг.
func (s *Service) GetBooks(ctx context.Context) (dto.Books, error) {
	booksSt, err := s.st.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	return dto.FromBooksSt(booksSt), nil
}

// GetAuthors возвращает список авторов.
func (s *Service) GetAuthors(ctx context.Context) (dto.Authors, error) {
	authorsSt, err := s.st.GetAuthors(ctx)
	if err != nil {
		return nil, err
	}

	return dto.FromAuthorsSt(authorsSt), nil
}

// GetBook возвращает книгу по идентификатору.
func (s *Service) GetBook(ctx context.Context, id int) (*dto.Book, error) {
	if id < 1 || id > math.MaxInt {
		errNew := berror.Validation().Obj("id")
		slog.Log(ctx, errNew.Level(), errNew.Error())
		return nil, errNew
	}

	bookSt, err := s.st.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}

	book := dto.FromBookSt(*bookSt)

	return &book, nil
}
