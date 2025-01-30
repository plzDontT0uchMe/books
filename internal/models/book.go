package models

// Book - модель книги хранящаяся в базе данных (в моём случае в памяти)

type Book struct {
	Id      int     `json:"id"`
	Authors Authors `json:"authors"`
	Title   string  `json:"title"`
	Year    int     `json:"year"`
}

type Books []Book
