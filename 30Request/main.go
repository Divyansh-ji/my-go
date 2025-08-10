package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// client vs normal get call

	resp, err := http.Get("http://example.com")

	if err != nil {
		fmt.Println("Error occured", err)

	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("The readed text is", string(body))
	//using http client used to send the request

	//Why we use client??

	//http.client is acts as browser for sending http request and handling response
	//we use client instead of basic get bcoz it allow to send customize requests
	// like we set headers timeout etc stuffs

	client := &http.Client{} //we made our client
	//respp, err := client.Get("http://example.com")
	//making request with custome client

	if err != nil {
		fmt.Println("Error", err)
	}
	//bodyy, _ := io.ReadAll(respp.Body)
	//fmt.Println("readed text 2", string(bodyy))

	// both the things gives me exact output
	// If I want to customize the request we have to use of http.NewRequest

	respp, _ := http.NewRequest("GET", "http://example.com", nil)
	respp.Header.Add("User-Agent", "MyCustomeClient/1.0 ")
	respp.Header.Add("Authorization", "Bearer mytoke")

	req, _ := client.Do(respp) //this is will send the request

	bodyy, _ := io.ReadAll(req.Body)
	fmt.Println("readed text", string(bodyy))
	defer req.Body.Close()

	//Type response example

	resp, err = http.Get("https://example.com/data")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.StatusCode, resp.Status)
	content, _ := io.ReadAll(resp.Body)
	fmt.Println("Body length:", resp.ContentLength)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Body content:", string(content))

	//String way to make request

}
