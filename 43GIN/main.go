package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //this line say gin give me a server

	// Simple health-check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"message": "hello" + name,
		})
	})
	r.GET("/profile", retjson)
	r.GET("/add", qwithgin)

	r.GET("/pongo", func(c *gin.Context) {
		time.Sleep(2 * time.Second) // simulate some processing
		c.JSON(200, gin.H{"message": "pong"})
	})
	//r.Use(TimmingMiddleware())

	//private method for authentication stuff

	private := r.Group("/private")
	private.Use(AuthMiddleware())
	{
		private.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "This is private data"})
		})
	}

	r.Run(":8080") // Run server on port 8080

}

func qwithgin(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")

	//converting string to integer

	xInt, _ := strconv.Atoi(x)
	yInt, _ := strconv.Atoi(y)

	c.JSON(200, gin.H{
		"sum": xInt + yInt,
	})
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func retjson(c *gin.Context) {
	user := User{ID: 1, Name: "Divyansh", Email: "divyanshtiwary@babu.com"}
	c.JSON(200, user)
}
func loggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		fmt.Println("Incoming request", method, path)

		ctx.Next() // go to next middleware /handler

	}
}

func TimmingMiddleware() gin.HandlerFunc {
	return func(sher *gin.Context) {
		start := time.Now()

		sher.Next() //procees the requeset

		duration := time.Since(start)
		fmt.Println("Request duration time is  ", duration)

	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-api-key")
		if apiKey != "secret123" {
			ctx.JSON(401, gin.H{
				"error": "unauthorised",
			})
			ctx.Abort()
			return

		}
		ctx.Next()
	}
}
