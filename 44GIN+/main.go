package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/search", func(c *gin.Context) {
		query := c.Query("query")              // get query param "query"
		limit := c.DefaultQuery("limit", "10") // get "limit", default = 10 if not provided

		c.JSON(http.StatusOK, gin.H{
			"query": query,
			"limit": limit,
		})
	})
	r.GET("/product", productsss)
	r.GET("/users/:id", pathpp)
	r.GET("/product/:title", pathpx)
	r.GET("/productt/:idd", pathquery)
	r.GET("/login", loginhandler)
	r.POST("/register", praclogin)

	r.Run(":8080")
}

func productsss(ctx *gin.Context) {

	query := ctx.Query("category")
	limit := ctx.DefaultQuery("sort", "asc")
	ctx.JSON(200, gin.H{
		"category": query,
		"sort":     limit,
	})

}

// Reading path parameters
func pathpp(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{
		"user_id": id,
	})

}

//practise

func pathpx(ctx *gin.Context) {
	title := ctx.Param("title")
	ctx.JSON(200, gin.H{
		"books": title,
	})
}

func pathquery(c *gin.Context) {
	//path parameter
	idd := c.Param("123")
	//query parameter

	queryy := c.DefaultQuery("USD", "INR")

	c.JSON(200, gin.H{
		"product id": idd,
		"currency":   queryy,
	})

}

// reading json
func loginhandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	remember := c.DefaultQuery("remember", "false")

	c.JSON(200, gin.H{
		"username": username,
		"password": password,
		"remember": remember,
	})
}

//practise

func praclogin(c *gin.Context) {
	name := c.PostForm("name")   // feild name
	email := c.PostForm("email") // feild name
	age := c.PostForm("age")     // feild name

	c.JSON(200, gin.H{
		"name":  name,
		"email": email,
		"age":   age,
	})

}
