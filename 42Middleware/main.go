package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Task 1:
//Create a middleware that:
//Prints "Request received" before the handler runs.
//Prints "Response sent" after the handler runs.
//Attach it globally with r.Use()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside handler :Home")
}

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/home", homeHandler).Methods("GET")
	r.Use(Logmiddleware)

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

//understanding logging and timming

func Logmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, r.RequestURI)
		p := time.Now()
		next.ServeHTTP(w, r)
		e := time.Since(p)

		log.Printf("the time take is ", e)

	})
}
