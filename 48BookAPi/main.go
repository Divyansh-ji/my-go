package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json :"id"`
	Title  string `json :"title"`
	Author string `json :"author"`
	Year   int    `json:"year"`
}

var books []Book

func main() {

	r := gin.Default()
	r.POST("/book", postbook)
	r.GET("/book", listbook)
	r.GET("/book/query", bookbyaut)
	r.Run(":9090")

}

func postbook(c *gin.Context) {

	var book Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	books = append(books, book)
	c.JSON(200, gin.H{
		"message": "book added successfully",
		"book":    book,
	})

}
func listbook(c *gin.Context) {
	c.JSON(200, gin.H{
		"print the books": books,
	})
}

// searching book by author
func bookbyaut(c *gin.Context) {
	aut := c.Query("author")
	var result []Book

	for _, booky := range books {
		if aut == booky.Author {
			result = append(result, booky)
		}
	}
	if len(result) == 0 {
		c.JSON(404, gin.H{
			"error": "no books found for this author",
		})
		return
	}
	c.JSON(200, result)

}
