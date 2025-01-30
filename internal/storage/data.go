package storage

import "backend/go/books/internal/models"

// Хранение данных в памяти

var Authors = models.Authors{
	{
		Id:   1,
		Name: "Leo Tolstoy",
	},
	{
		Id:   2,
		Name: "Fyodor Dostoevsky",
	},
	{
		Id:   3,
		Name: "Anton Chekhov",
	},
	{
		Id:   4,
		Name: "Mark Twain",
	},
	{
		Id:   5,
		Name: "Charles Dickens",
	},
}

var Books = models.Books{
	{
		Id:      1,
		Authors: Authors[0:1],
		Title:   "War and Peace",
		Year:    1869,
	},
	{
		Id:      2,
		Authors: Authors[1:2],
		Title:   "Crime and Punishment",
		Year:    1866,
	},
	{
		Id:      3,
		Authors: Authors[2:3],
		Title:   "The Cherry Orchard",
		Year:    1904,
	},
	{
		Id:      4,
		Authors: Authors[3:4],
		Title:   "Adventures of Huckleberry Finn",
		Year:    1884,
	},
	{
		Id:      5,
		Authors: Authors[4:5],
		Title:   "A Tale of Two Cities",
		Year:    1859,
	},
	{
		Id:      6,
		Authors: Authors[0:2],
		Title:   "Collaborative Work",
		Year:    2023,
	},
}
