package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var notes []Note // it will store the many notes as it is the form of slice

func main() {

	r := gin.Default()
	r.POST("/notes", listNotes)        // Create
	r.GET("/notes", getNotes)          // Read all
	r.GET("/notes/:id", getbyId)       // Read one
	r.PUT("/notes/:id", updatebyid)    // Update
	r.DELETE("/notes/:id", deletebyid) // Delete

	r.Run(":7070")
}
func listNotes(c *gin.Context) {
	var note Note // this only store one note we

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	notes = append(notes, note)
	// router hit kiya naya json aya uske andar naya notes usko struct mein over ride karaya note wale fir uskko slice mein add karliya

	c.JSON(http.StatusCreated, gin.H{
		"message": "Note created successfully",
		"note":    note,
	})

}
func getNotes(c *gin.Context) {
	c.JSON(http.StatusOK, notes)

}
func getbyId(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	for _, note := range notes {
		if id == note.ID {
			c.JSON(200, note)
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

	}
}
func updatebyid(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	var update Note // new data we need to update will store here

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	// at which id we need to update
	for i, uouo := range notes {
		if id == uouo.ID {
			notes[i].Title = update.Title
			notes[i].Content = update.Content
			c.JSON(200, notes[i])
			return

		}
	}
	c.JSON(404, gin.H{
		"error": "note not found",
	})

}

// Deleteing by id
func deletebyid(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}

	for i, del := range notes {
		if id == del.ID {
			notes = append(notes[:i], notes[i+1:]...)
			c.JSON(200, gin.H{
				"message": "note deleted succesufll",
			})
			return

		}
	}

}
