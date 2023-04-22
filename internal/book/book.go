package book

import "github.com/deoraczo/book-manager-cli/internal/book/dto"

type Book struct {
	ID        int64
	Title     string
	Author    string
	Publisher string
	Year      int
}

var Books []Book
var nextID int64 = 1

func FromDTO(data dto.CreateBookDTO) Book {
	id := nextID
	nextID++

	return Book{
		ID:        id,
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		Year:      data.Year,
	}
}
