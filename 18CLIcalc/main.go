package main

import "fmt"

func main() {

	var a, b float64
	var op string
	fmt.Println("enter the number 1 ")
	fmt.Scan(&a)
	fmt.Println("enter the operation(+ , - ,/ , *)")
	fmt.Scan(&op)
	fmt.Println("Enter the number 2")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a - b)
	case "/":
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("error")
		}
	default:
		fmt.Println("hey baba")

	}

}
