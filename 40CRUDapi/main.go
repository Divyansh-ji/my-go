package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: 101, Name: "Divyansh"},
	{ID: 102, Name: "Harsh"},
	{ID: 103, Name: "Anu"},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Get/{name}", Getitembyname).Methods("GET")
	http.ListenAndServe(":9999", r)
}

func Getitembyname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchName := vars["name"]

	for _, item := range items {
		if item.Name == searchName {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return // stop here after sending the response
		}
	}

	// If no match found
	http.Error(w, "Item not found", http.StatusNotFound)
}
