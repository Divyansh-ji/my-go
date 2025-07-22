package main

import "fmt"

func main() {
	fmt.Println("if and else in golang")
	logincount := 23
	var result string

	if logincount < 10 {
		result = " Regular user"
	} else if logincount > 10 {
		result = "Watch out "
	} else {
		result = "exactly 10 user "
	}

	fmt.Println(result)

	//special syntax

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
	if num := 9; num < 0 {
		fmt.Println("number is negative ")
	} else if num < 10 {
		fmt.Println(num, "has  1 digit")
	} else {
		fmt.Println(num, "has multiple digit ")
	}

}
