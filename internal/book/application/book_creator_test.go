package application

import (
	"testing"

	"github.com/deoraczo/book-manager-cli/internal/book"
	"github.com/deoraczo/book-manager-cli/internal/book/dto"
)

func TestBookCreator_Execute(t *testing.T) {
	bookCreator := NewBookCreator()
	bookDto := dto.CreateBookDTO{
		Title:     "Tests",
		Author:    "Some",
		Publisher: "Jojo",
		Year:      2021,
	}

	err := bookCreator.Execute(bookDto)

	if err != nil {
		t.Errorf("Error creating book: %v", err)
	}

	if len(book.Books) != 1 {
		t.Errorf("Expected 1 book, got %v", len(book.Books))
	}

	createdBook := book.Books[0]

	if createdBook.Title != bookDto.Title {
		t.Errorf("Expected title %v, got %v", bookDto.Title, createdBook.Title)
	}

	if createdBook.Publisher != bookDto.Publisher {
		t.Errorf("Expected publisher %v, got %v", bookDto.Publisher, createdBook.Publisher)
	}

	if createdBook.Author != bookDto.Author {
		t.Errorf("Expected author %v, got %v", bookDto.Author, createdBook.Author)
	}

	if createdBook.Year != bookDto.Year {
		t.Errorf("Expected year %v, got %v", bookDto.Year, createdBook.Year)
	}
}
