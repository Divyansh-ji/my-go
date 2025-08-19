package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json :"id"`
	Title       string `json:"title"`
	Description string `json :"description"`
	Duedate     int    `json : "due date"`
}

var tasks []Task
var nextID = 1

func main() {

	r := gin.Default()

	r.POST("/task", gettingtask)

	r.Run(":9000")

}
func addTask(t Task) Task {
	t.ID = nextID // assign current ID
	nextID++
	tasks = append(tasks, t)
	return t
}
func gettingtask(ctx *gin.Context) {

	var task Task

	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	nextTask := addTask(task)
	ctx.JSON(http.StatusCreated, nextTask)

}
func updatedTask(ctx *gin.Context) {
	params := ctx.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	var update Task

	if err := ctx.BindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Title = update.Title
			tasks[i].Description = update.Description
			tasks[i].Duedate = update.Duedate
			ctx.JSON(200, tasks[i])
			return

		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}
