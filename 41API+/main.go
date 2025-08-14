package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet", greeting).Methods("GET")
	r.HandleFunc("/square", squu).Methods("GET")
	r.HandleFunc("/posts", postby).Methods("GET")
	r.HandleFunc("/filter", filterby).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func squu(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	numStr := params.Get("num")

	if numStr == "" {
		http.Error(w, "Missing 'num' parameter", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	result := num * num
	fmt.Fprintf(w, "%d", result)
}
func postby(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	user := params.Get("userID")

	if user == "" {
		http.Error(w, "userid query parameter is required", http.StatusBadRequest)
		return
	}
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=<userId>")

	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusBadRequest)
		return

	}
	defer resp.Body.Close()

	var posts []Post

	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		http.Error(w, "Failed to parse the post", http.StatusBadRequest)

	}
	for _, post := range posts {
		fmt.Println(w, post.Title)

	}
}

func filterby(w http.ResponseWriter, r *http.Request) {
	var list []int
	params := r.URL.Query()
	minn := params.Get("min")
	minnum, err := strconv.Atoi(minn)
	if err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
	}

	maxx := params.Get("max")
	maxnum, err := strconv.Atoi(maxx)
	if err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
	}

	for i := minnum; i <= maxnum; i++ {

		list = append(list, i)

	}
	fmt.Fprintf(w, "so the list is  %v", list)

}
