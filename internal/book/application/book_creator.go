package application

import (
	"github.com/deoraczo/book-manager-cli/internal/book"
	"github.com/deoraczo/book-manager-cli/internal/book/dto"
)

type BookCreator struct {
}

func NewBookCreator() BookCreator {
	return BookCreator{}
}

func (b BookCreator) Execute(data dto.CreateBookDTO) error {

	book.Books = append(book.Books, book.FromDTO(data))

	return nil
}
