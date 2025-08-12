package main

import (
	"encoding/json"
	"net/http"
	"strings"

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

	// Loop and check for a match (case-insensitive)
	for _, item := range items {
		if strings.EqualFold(item.Name, searchName) { // EqualFold ignores case
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// If no match found
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
}
