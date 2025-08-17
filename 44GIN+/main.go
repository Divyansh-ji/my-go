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
