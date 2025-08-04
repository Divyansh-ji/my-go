package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Step 1: Make the GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	// Step 2: Always close the response body
	defer resp.Body.Close()

	// Step 3: Check status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	// Step 4: Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body:", err)
	}

	// Step 5: Print response
	fmt.Println(string(body))

	//Custom header with GET in go
	//request has been created

	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error occurred")
	}

	//This is the step to set custom headder

	req.Header.Set("User-Agent", "G0-client/1.0")
	req.Header.Set("Authorization", "Bearer my-secret-token")

	//using http client used to send the request

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error in sending requenst", err)
	}
	defer resp.Body.Close()

	//Now will read the response

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in sending requenst", err)

	}
	fmt.Println("Status", resp.Status)
	fmt.Println("Response ", string(body))

}
