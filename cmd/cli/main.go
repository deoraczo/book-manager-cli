package main

import (
	"fmt"

	"github.com/deoraczo/book-manager-cli/internal/book/application"
	"github.com/deoraczo/book-manager-cli/internal/book/dto"
)

type menuOption int

const (
	exit menuOption = iota
	createBook
	getBooks
)

func main() {
	bookCreator := application.NewBookCreator()
	bookRetriever := application.NewBookRetriever()

	for {
		printMenu()

		option := readOption()

		switch option {
		case createBook:
			bookDto := readBookData()

			handleCreateBook(bookCreator, bookDto)
			fmt.Println("Libro creado correctamente.")
		case getBooks:
			books, _ := bookRetriever.Execute()
			fmt.Println(books)
		case exit:
			fmt.Println("Adiós!")
			return
		default:
			fmt.Println("Opción inválida, por favor intente nuevamente.")
		}
	}
}

// Imprimir las opciones del menú
func printMenu() {
	fmt.Println("Seleccione una opción:")
	fmt.Printf("%v: Crear libro\n", createBook)
	fmt.Printf("%v: Mostrar todo los libros\n", getBooks)
	fmt.Printf("%v: Salir\n", exit)
}

// Leer la opción seleccionada por el usuario
func readOption() menuOption {
	var option menuOption
	fmt.Scanln(&option)
	return option
}

// Leer los datos de un nuevo libro
func readBookData() dto.CreateBookDTO {
	var book dto.CreateBookDTO

	fmt.Println("Ingrese el título:")
	fmt.Scanln(&book.Title)

	fmt.Println("Ingrese el autor:")
	fmt.Scanln(&book.Author)

	fmt.Println("Ingrese el editor:")
	fmt.Scanln(&book.Publisher)

	fmt.Println("Ingrese el año:")
	fmt.Scanln(&book.Year)

	return book
}

func handleCreateBook(c application.BookCreator, d dto.CreateBookDTO) {
	c.Execute(d)
}
