package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []Task
var nextID int = 1

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var choice int

	for {
		fmt.Println("\n=== Todo-CLI ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Addtask(reader)
		case 2:
			ListTask()
		case 3:
			completetask(reader)
		case 4:
			deletetask(reader)
		case 5:
			fmt.Println("👋 Exiting...")
			return
		default:
			fmt.Println("❌ Invalid choice")
		}
	}
}

func Addtask(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	tasks = append(tasks, Task{ID: nextID, Title: task, Done: false})
	nextID++
	fmt.Println("✅ Task added!")
}

func ListTask() {
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found.")
		return
	}
	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", t.ID, t.Title, status)
	}
}

func completetask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to mark complete: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			fmt.Println("✔ Task marked as complete:", tasks[i].Title)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("❌ Task ID not found.")
	}
}

func deletetask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	index := -1
	for i := range tasks {
		if tasks[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("❌ Task ID not found.")
		return
	}

	fmt.Println("🗑 Deleted:", tasks[index].Title)
	tasks = append(tasks[:index], tasks[index+1:]...)
}
