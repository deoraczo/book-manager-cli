package application

import "github.com/deoraczo/book-manager-cli/internal/book"

type BookRetriever struct {
}

func NewBookRetriever() BookRetriever {
	return BookRetriever{}
}

func (r BookRetriever) Execute() ([]book.Book, error) {
	// Crear una nueva slice para contener los mismos elementos que la original
	booksCopy := make([]book.Book, len(book.Books))
	copy(booksCopy, book.Books)

	// Retornar la nueva slice
	return booksCopy, nil
}
