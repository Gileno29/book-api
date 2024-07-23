package main

import (
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

	router.Run("localhost:8080")

}
