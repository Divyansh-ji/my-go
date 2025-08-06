package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	//GET with query Parameters

	baseURl := "https://jsonplaceholder.typicode.com/posts"
	u, _ := url.Parse(baseURl)

	//add query params ? userID=1

	params := u.Query()
	params.Set("UserId", "1") //add as a map key-val pair
	u.RawQuery = params.Encode()
	fmt.Println("Requesting Url", u.String())

	resp, err := http.Get(u.String())

	if err != nil {
		fmt.Println("Error getting ", u.String())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
