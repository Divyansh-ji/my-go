package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Post struct {
	ID    int
	Title string
	Body  string
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error fetching post:", err)
		return
	}
	defer resp.Body.Close()

	var post Post

	//This reads JSON directly from the response body (no need to read all first)
	json.NewDecoder(resp.Body).Decode(&post)

	fmt.Println("Post ID:", post.ID)
	fmt.Println("Title:", post.Title)

	//for encoder
	postt := Post{ID: 1, Title: "mai hoon Tiger", Body: "salman bhai jaise"}

	json.NewEncoder(os.Stdout).Encode(postt)
}
