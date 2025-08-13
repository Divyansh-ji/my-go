package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet", greeting).Methods("GET")
	r.HandleFunc("/square", squu).Methods("GET")
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
