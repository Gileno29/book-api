package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Quatity int    `json:"quality"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quatity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott", Quatity: 2},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quatity: 3},
}

// Funcao para retornar uma lista de books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Percorre as struct de books e encontra o livro pelo ID
func getBooksById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")

}

// Funcao para tratamento da rota de buscar Livro pelo ID
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBooksById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// Funcao para criar um novo livro
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	//pega uma instlncia do GIN para manipular as rotas
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)

	router.Run("localhost:8080")

}
