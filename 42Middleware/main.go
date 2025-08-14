package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside handler :Home")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.Use(middleware)
	log.Println("Server running on:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		//this will call the actual handlee that is thwe homehandler
		next.ServeHTTP(w, r)
		fmt.Println("Response sent")
		next.ServeHTTP(w, r)
	})
}
