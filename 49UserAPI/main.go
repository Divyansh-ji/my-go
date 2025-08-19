package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email" binding: "required email"`
	Age   int    `json: age binding: gte=13,lte=100`
}

var users []User // suppose this contain many users

func main() {

	r := gin.Default()

	r.POST("/users", newUser)
	r.GET("/users:id", getbyId)
	r.GET("/user:id", getalltheid)

}
func newUser(c *gin.Context) {
	var newu User

	if err := c.BindJSON(&newu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// added to orginal databases

	users = append(users, newu)
	c.JSON(200, gin.H{
		"added succesfully ": users,
	})

}
func getbyId(c *gin.Context) {

	id := c.Param("id")
	idstr, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	var result []User

	for _, t := range users {
		if idstr == t.ID {
			result = append(result, t)
			c.JSON(200, gin.H{
				"resultant user": result,
			})

		}
		if len(result) == 0 {
			c.JSON(404, gin.H{
				"error": "no books found for this author",
			})
			return
		}

	}
	c.JSON(200, result)

}
func getalltheid(c *gin.Context) {
	c.JSON(200, gin.H{
		"messsage": users,
	})

}
