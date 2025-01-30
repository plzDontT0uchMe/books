package dto

import "backend/go/books/internal/models"

// Book структура для книги (DTO, в данном случае возможно избыточно).

type Book struct {
	Id      int     `json:"id"`
	Authors Authors `json:"authors"`
	Title   string  `json:"title"`
	Year    int     `json:"year"`
}

func (b *Book) ToSt() models.Book {
	return models.Book{
		Id:      b.Id,
		Authors: b.Authors.ToSt(),
		Title:   b.Title,
		Year:    b.Year,
	}
}

func FromBookSt(b models.Book) Book {
	return Book{
		Id:      b.Id,
		Authors: FromAuthorsSt(b.Authors),
		Title:   b.Title,
		Year:    b.Year,
	}
}

type Books []Book

func (b *Books) ToSt() models.Books {
	var books models.Books
	for _, book := range *b {
		books = append(books, book.ToSt())
	}
	return books
}

func FromBooksSt(books models.Books) Books {
	var b Books
	for _, book := range books {
		b = append(b, FromBookSt(book))
	}
	return b
}
