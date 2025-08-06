package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// there are serveral cases arises in encoding and decoding of JSON

	//we can modifed are struct when we are decoding json  for only required we only used reequired feild

	//GENERIC JSON with interface

	//For dynamic json we use interface when we dont know how many feild are to be used are what are the feild we will decode form json

	//example

	jsonData := []byte(`{
	   "id": 1,
        "name": "Divyansh Tiwari",
        "skills": ["Go", "Docker", "Kubernetes"],
        "isActive": true

	}`)

	//declared empty interface to hold the decoded data

	var result interface{}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		panic(err)
	}
	// Type assertion : convert interface{} to map[string]interface{}

	data := result.(map[string]interface{})

	// Access the feil dynamically

	fmt.Println("ID", data["id"])
	fmt.Println("Name", data["name"])
	fmt.Println("Skills", data["skills"])
	fmt.Println("Is Active", data["isActive"])

	// lets loop into the map

	fmt.Println("\n-- All Key-Value Pairs --")
	for key, value := range data {
		fmt.Println("%s : %v\n", key, value)
	}

}
