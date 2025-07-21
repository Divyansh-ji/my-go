package main

import "fmt"

func main() {
	fmt.Println("weclom to golangs map")
	language := make(map[int]string)

	language[1] = "javascript"
	language[2] = "Ruby"
	language[3] = "python"
	fmt.Println("language as a map", language)
	fmt.Println("1 shorts for ", language[1])
	delete(language, 1)
	fmt.Println("List after delete ", language)

	// loops are interesting in golang

	for key, value := range language {
		fmt.Printf("for key %v , value is %v\n", key, value)
	}

}
