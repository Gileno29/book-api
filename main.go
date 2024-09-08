package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type book struct {
	ID      string
	Title   string
	Author  string
	Quatity int
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

func chockoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID query parameter"})
		return
	}

	book, err := getBooksById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	if book.Quatity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quatity -= 1

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID query parameter"})
		return
	}

	book, err := getBooksById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	if book.Quatity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quatity += 1

	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	//pega uma instlncia do GIN para manipular as rotas
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", chockoutBook)
	router.PATCH("/return", returnBook)
	router.Run("0.0.0.0:8080")

}
