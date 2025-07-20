package main

import (
	"fmt"
)

const Logintoken string = "this is live"

func main() {
	var username string = "Divyansh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.569483947329
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVaribale int
	fmt.Println(anotherVaribale)
	fmt.Printf("Variable is of type: %T \n", anotherVaribale)

	// implicit type
	var website = "instagram.com"
	fmt.Println(website)

	// no var style

	numberofUser := 300000.0
	fmt.Println(numberofUser)
	//only allowed inside the method

	fmt.Println(Logintoken)
	fmt.Printf("Variable is of type: %T \n", Logintoken)

}
