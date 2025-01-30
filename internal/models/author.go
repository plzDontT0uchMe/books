package models

// Author - модель книги хранящаяся в базе данных (в моём случае в памяти)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Authors []Author
