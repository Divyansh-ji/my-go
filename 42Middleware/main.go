package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Task 1:
// Create a middleware that:
// Prints "Request received" before the handler runs.
// Prints "Response sent" after the handler runs.
// Attach it globally with r.Use()
func publicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the public route!"))
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the private route!"))
}

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/", homeHandler).Methods("GET")

	r.Use(loggingmiddleware)

	r.HandleFunc("/public", publicHandler).Methods("GET")
	// private subrouter (logging + authentication)

	private := r.PathPrefix("/private").Subrouter()
	private.Use(authenticationmiddleware)
	private.HandleFunc("", privateHandler).Methods("GET")
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)

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

//Task 3:
//Write an authentication middleware that:
//• Checks for a header X-API-KEY with the value "mysecret".
//• If missing or incorrect → return http.StatusForbidden .
//• If correct → allow the request to

type authenticationmiddlewaree struct {
	tokenUser map[string]string
}

// making
func (amw *authenticationmiddlewaree) populate() {
	amw.tokenUser["000000"] = "user0"
	amw.tokenUser["aaaaaa"] = "userA"
	amw.tokenUser["X-API-KEY"] = "mysecret"

}
func (amw *authenticationmiddlewaree) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-KEY")

		if user, found := amw.tokenUser[token]; found {
			log.Printf("Authenticated user %s\n", user)

			//since we have found the correct request so let the user login so will pass the to the handle functions so

			next.ServeHTTP(w, r)

		} else {
			http.Error(w, "forbidden", http.StatusForbidden)
		}

	})

}

//Level 4 - Combine Multiple Middlewares
//Task 4:
//• Create:
//• Add both the logging and authentication middlewares to your router.
//• /public route → only logging middleware applies.
//• /private route → logging + authentication applies.//
//Hint: Use Subrouter() for /private.
//Goal: Practice global vs. route-specific

func loggingmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})

}
func authenticationmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-API-KEY")
		if token != "mysecret" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		log.Println("Authenticated successfully")
		next.ServeHTTP(w, r)

	})
}
