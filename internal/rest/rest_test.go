package rest_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/go/books/internal/config"
	"backend/go/books/internal/dto"
	e "backend/go/books/internal/error"
	"backend/go/books/internal/models"
	"backend/go/books/internal/rest"
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

func initServer(t *testing.T) (http.Server, *mock_interfaces.MockStorage) {
	ctrl := gomock.NewController(t)
	cfg := config.Server{
		Port:              8080,
		RateLimitCount:    100,
		RateLimitDuration: 60,
		RequestTimeout:    time.Second * 30,
	}
	st := mock_interfaces.NewMockStorage(ctrl)
	sv := service.New(st)
	return rest.New(cfg, sv), st
}

func TestBooksHandler_Success(t *testing.T) {
	server, st := initServer(t)

	st.EXPECT().GetBooks(gomock.Any()).Return(booksSt, nil)

	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var books dto.Books
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &books); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, books.ToSt(), booksSt)
}

func TestBooksHandler_Fail(t *testing.T) {
	server, st := initServer(t)

	errExpected := berror.Internal()

	st.EXPECT().GetBooks(gomock.Any()).Return(nil, fmt.Errorf("error"))

	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Fatalf("expected status %d, got %d", resp.Code, http.StatusOK)
	}

	var errBase berror.Base
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &errBase); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, true, e.Equal(errBase, errExpected))
}

func TestAuthorsHandler_Success(t *testing.T) {
	server, st := initServer(t)

	st.EXPECT().GetAuthors(gomock.Any()).Return(authorsSt, nil)

	req, err := http.NewRequest(http.MethodGet, "/authors", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var authors dto.Authors
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &authors); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, authors.ToSt(), authorsSt)
}

func TestAuthorsHandler_Fail(t *testing.T) {
	server, st := initServer(t)

	errExpected := berror.Internal()

	st.EXPECT().GetAuthors(gomock.Any()).Return(nil, fmt.Errorf("error"))

	req, err := http.NewRequest(http.MethodGet, "/authors", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Fatalf("expected status %d, got %d", resp.Code, http.StatusOK)
	}

	var errBase berror.Base
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &errBase); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, true, e.Equal(errBase, errExpected))
}

func TestBookHandler_Success(t *testing.T) {
	server, st := initServer(t)

	st.EXPECT().GetBook(gomock.Any(), 1).Return(&booksSt[0], nil)

	req, err := http.NewRequest(http.MethodGet, "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var book dto.Book
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &book); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, book.ToSt(), booksSt[0])
}

func TestBookHandler_Fail(t *testing.T) {
	server, st := initServer(t)

	errExpected := berror.Internal()

	st.EXPECT().GetBook(gomock.Any(), 1).Return(nil, fmt.Errorf("error"))

	req, err := http.NewRequest(http.MethodGet, "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Fatalf("expected status %d, got %d", resp.Code, http.StatusOK)
	}

	var errBase berror.Base
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &errBase); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, true, e.Equal(errBase, errExpected))
}

func TestBookHandler_Validation(t *testing.T) {
	server, _ := initServer(t)

	errExpected := berror.Validation().Obj("id")

	req, err := http.NewRequest(http.MethodGet, "/book/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Fatalf("expected status %d, got %d", resp.Code, http.StatusOK)
	}

	var errBase berror.Base
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &errBase); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, true, e.Equal(errBase, errExpected))
}

func TestBookHandler_Invalid(t *testing.T) {
	server, _ := initServer(t)

	errExpected := berror.InvalidArgument().Obj("id")

	req, err := http.NewRequest(http.MethodGet, "/book/asd", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Fatalf("expected status %d, got %d", resp.Code, http.StatusOK)
	}

	var errBase berror.Base
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &errBase); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, true, e.Equal(errBase, errExpected))
}

func TestBookHandler_NotFound(t *testing.T) {
	server, st := initServer(t)

	errExpected := berror.NotFound().Obj("book")

	st.EXPECT().GetBook(gomock.Any(), 1).Return(nil, errExpected)

	req, err := http.NewRequest(http.MethodGet, "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Code == http.StatusOK {
		t.Fatalf("expected status %d, got %d", resp.Code, http.StatusOK)
	}

	var errBase berror.Base
	if errUnmarshal := json.Unmarshal(resp.Body.Bytes(), &errBase); errUnmarshal != nil {
		t.Fatal(errUnmarshal)
	}

	assert.Equal(t, true, e.Equal(errBase, errExpected))
}
