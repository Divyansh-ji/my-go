package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	//JSON handling

	//What is json?
	//JSON (JavaScript Object Notation) is a simple data interchange format. Syntactically it resembles the objects and lists of JavaScript. It is most commonly used for communication between web back-ends and JavaScript programs running in the browser, but it is used in many other places, too

	// with json package its too easy to read and write Json data form your GO program

	// json provide two function to do that marhal and unmarshal

	//conversion :

	//Json -> struct(go data structure ) using unmarshal
	//struct -> JSON(universal formate) using marshal
	//

	// Example to understand

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error has been occured", err)
	}
	defer resp.Body.Close()

	//we read response to check the feild we need to write in struct so we can get all the data from json\

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Readed data", string(body))

	var J Post

	//decoding usinh unmarshal

	json.Unmarshal(body, &J) //&J we are telling that put in J struct we are pointing to J

	fmt.Println("UserID", J.UserID)
	fmt.Println("ID", J.ID)
	fmt.Println("Tittle", J.Title)

	// Example of marshal we will encode our data in json

	user := Post{
		ID:    1,
		Title: "Divyansh Tiwari",
		Body:  "divyansh@example.com",
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	//print the json string

	fmt.Println(string(jsonData))

}

// we will make this struct after reading the json
//we dont know which feild are present

type Post struct {
	UserID int    `json:"abdul"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
