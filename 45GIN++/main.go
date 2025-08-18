package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
type Book struct {
	Title     string   `json:"title"`
	Content   string   `json :"content"`
	Tags      []string `json: "tags"`
	Published bool     `json:publised`
}

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var user User

		// Bind JSON body into struct
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		})
	})

	r.POST("/article", func(ctx *gin.Context) {

		var book Book

		// binding the json body into the struct

		if err := ctx.BindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{
			"title":     book.Title,
			"content":   book.Content,
			"tags ":     book.Tags,
			"published": book.Published,
		})
	})

	r.POST("/valid", func(ctx *gin.Context) {

		var userr Userr
		//bind and validate
		if err := ctx.ShouldBind(&userr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "registration successfull",
			"user":    userr,
		})
	})

	r.Run(":8080")
}

type Userr struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"gte=18,lte=100"`
}
