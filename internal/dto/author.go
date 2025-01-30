package dto

import "backend/go/books/internal/models"

// Author структура для автора (DTO, в данном случае возможно избыточно).

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (a *Author) ToSt() models.Author {
	return models.Author{
		Id:   a.Id,
		Name: a.Name,
	}
}

func FromAuthorSt(authorSt models.Author) Author {
	return Author{
		Id:   authorSt.Id,
		Name: authorSt.Name,
	}
}

type Authors []Author

func (a Authors) ToSt() models.Authors {
	var authorsSt models.Authors
	for _, author := range a {
		authorsSt = append(authorsSt, author.ToSt())
	}
	return authorsSt
}

func FromAuthorsSt(authorsSt models.Authors) Authors {
	var authors Authors
	for _, a := range authorsSt {
		authors = append(authors, FromAuthorSt(a))
	}
	return authors
}
