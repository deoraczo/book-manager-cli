package application

import (
	"reflect"
	"testing"

	"github.com/deoraczo/book-manager-cli/internal/book"
)

func TestBookRetriever_Execute(t *testing.T) {
	loadBooksInMemory()
	retriever := NewBookRetriever()

	actualBooks, err := retriever.Execute()

	// Verificar que no hay error
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(book.Books, actualBooks) {
		t.Errorf("Actual books (%v) is not equal to expected books (%v)", actualBooks, book.Books)
	}
}

func TestBookRetrieverIsImmutable(t *testing.T) {
	loadBooksInMemory()
	retriever := NewBookRetriever()
	updatedTitle := "Don Quijote"

	actualBooks, err := retriever.Execute()

	// Verificar que no hay error
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	//intentamos modificar el slice de libros
	actualBooks[0].Title = updatedTitle

	// Comprobamos que la modificación no afectó al slice original
	/*if book.Books[0].Title == updatedTitle {
		t.Errorf("El título del primer libro debió mantenerse como 'El nombre del viento', pero es '%s'", book.Books[0].Title)
	}*/
	if book.Books[0].Title != "El nombre del viento" {
		t.Errorf("Expected 'El nombre del viento' but got '%s'", book.Books[0].Title)
	}
}

func loadBooksInMemory() {
	expectedBooks := []book.Book{
		{ID: 1, Title: "El nombre del viento", Author: "Patrick Rothfuss", Publisher: "Plaza & Janés", Year: 2007},
		{ID: 2, Title: "El temor de un hombre sabio", Author: "Patrick Rothfuss", Publisher: "Plaza & Janés", Year: 2011},
		{ID: 3, Title: "Harry Potter y la piedra filosofal", Author: "J.K. Rowling", Publisher: "Salamandra", Year: 1997},
		{ID: 4, Title: "Harry Potter y la cámara secreta", Author: "J.K. Rowling", Publisher: "Salamandra", Year: 1998},
		{ID: 5, Title: "Harry Potter y el prisionero de Azkaban", Author: "J.K. Rowling", Publisher: "Salamandra", Year: 1999},
	}

	book.Books = append(book.Books, expectedBooks...)
}
