package main

import (
	"bytes"

	"fmt"
	"net/http"
)

//lets handle client stuff in a crazy way
// Custom http.Client with timeout

//let understand the above stuff

//time out is basically use when we want to run our request for limited time

// old school way

func main() {
	// Create a custom HTTP client with a 2-second timeout
	//client := http.Client{
	//	Timeout: 2 * time.Second,
	//}

	// URL that takes longer than 2s to respond (for test, use slow API or simulate)
	//url := "https://httpstat.us/200?sleep=5000" // sleeps for 5 seconds

	// Make the GET request
	//resp, err := client.Get(url)
	//if err != nil {
	//	fmt.Println("Request failed:", err)
	//	return
	//}
	//defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)

	// modern way
	//Make a GET request to a slow API and handle timeout using context.

	//ctx, cancel := context.WithTimeout(context.Background(), 60000*time.Second)
	//defer cancel()

	//req, err := http.NewRequestWithContext(ctx, "GET", "https://httpstat.us/200?sleep=5000", nil)
	//if err != nil {
	//	panic(err)
	//}

	//respp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Println("Request failed:", err)
	//	return
	//}
	//defer respp.Body.Close()

	//fmt.Println("Response ", respp.Status)

	//TASK 2 Perform a POST request (send JSON data)

	jsonData := []byte(`{
	   "id": 1,
		"name": "Divyansh Tiwari",
		"skills": ["Go", "Docker", "Kubernetes"],
		"isActive": true

	}`)
	// No need to marshal jsonData, it's already a []byte
	// jsonnn, err := json.Marshal(jsonData)
	// if err != nil {
	// 	fmt.Println("Error is marshalling JSon:", err)
	// 	return
	// }

	// create a new http request
	reqq, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resppp, err := client.Do(reqq)
	if err != nil {
		panic(err)
	}
	defer resppp.Body.Close()
	fmt.Println("Response status", resppp.Status)
}
