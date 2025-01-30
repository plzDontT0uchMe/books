package service_test

import (
	"context"
	"fmt"
	"testing"

	e "backend/go/books/internal/error"
	"backend/go/books/internal/models"
	"backend/go/books/internal/service"
	"backend/go/books/internal/storage/mocks"
	"backend/go/books/pkg/berror"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var authorsSt = models.Authors{
	{Id: 1, Name: "Author1"},
	{Id: 2, Name: "Author2"},
	{Id: 3, Name: "Author3"},
}

var booksSt = models.Books{
	{Id: 1, Authors: authorsSt[0:1], Title: "Book1", Year: 123},
	{Id: 2, Authors: authorsSt[1:2], Title: "Book2", Year: 234},
	{Id: 3, Authors: authorsSt[2:3], Title: "Book3", Year: 345},
}

func initService(t *testing.T) (*service.Service, *mock_interfaces.MockStorage) {
	ctrl := gomock.NewController(t)
	st := mock_interfaces.NewMockStorage(ctrl)
	return service.New(st), st
}

func TestGetBooks_Success(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetBooks(gomock.Any()).Return(booksSt, nil)

	books, err := sv.GetBooks(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, books.ToSt(), booksSt)
	assert.NotEmpty(t, books)
}

func TestGetBooks_SuccessEmpty(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetBooks(gomock.Any()).Return(models.Books{}, nil)

	books, err := sv.GetBooks(context.Background())

	assert.NoError(t, err)
	assert.Empty(t, books)
}

func TestGetBooks_Fail(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetBooks(gomock.Any()).Return(models.Books{}, fmt.Errorf("error"))

	books, err := sv.GetBooks(context.Background())

	assert.Error(t, err)
	assert.Empty(t, books)
}

func TestGetAuthors_Success(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetAuthors(gomock.Any()).Return(authorsSt, nil)

	authors, err := sv.GetAuthors(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, authors.ToSt(), authorsSt)
	assert.NotEmpty(t, authors)
}

func TestGetAuthors_SuccessEmpty(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetAuthors(gomock.Any()).Return(models.Authors{}, nil)

	authors, err := sv.GetAuthors(context.Background())

	assert.NoError(t, err)
	assert.Empty(t, authors)
}

func TestGetAuthors_Fail(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetAuthors(gomock.Any()).Return(models.Authors{}, fmt.Errorf("error"))

	authors, err := sv.GetAuthors(context.Background())

	assert.Error(t, err)
	assert.Empty(t, authors)
}

func TestGetBook_Success(t *testing.T) {
	sv, st := initService(t)

	st.EXPECT().GetBook(gomock.Any(), 1).Return(&booksSt[0], nil)

	book, err := sv.GetBook(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, book.ToSt(), booksSt[0])
	assert.NotNil(t, book)
}

func TestGetBook_NotFound(t *testing.T) {
	sv, st := initService(t)

	errExpected := berror.NotFound().Obj("book")

	st.EXPECT().GetBook(gomock.Any(), 1).Return(nil, errExpected)

	book, err := sv.GetBook(context.Background(), 1)

	assert.Equal(t, true, e.Equal(err, errExpected))
	assert.Nil(t, book)
}

func TestGetBook_Validation(t *testing.T) {
	sv, _ := initService(t)

	errExpected := berror.Validation().Obj("id")

	book, err := sv.GetBook(context.Background(), 0)

	assert.Equal(t, true, e.Equal(err, errExpected))
	assert.Nil(t, book)
}
