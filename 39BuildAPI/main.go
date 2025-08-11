package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:message`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(Message{Message: "hellow world!"})
	})
	//Post endpoint

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}

		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)

		w.Header().Set("content-type", "application/json")

		json.NewEncoder(w).Encode(msg)

	})
	http.ListenAndServe(":8080", nil)

}
