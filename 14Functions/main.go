package main

import "fmt"

// we dont need to call main because it acts as a entry point for your program
func main() {

	fmt.Println("Welcome to functions in golang ")
	greeter()
	fmt.Println("Result is", result(89, 97))

	adder, mymessage := proAdder(2, 2, 3, 4, 4, 5, 6, 7, 8, 9, 0, 23)
	fmt.Println("ALl the sum is ", adder)
	fmt.Println(mymessage)

}
func greeter() {
	fmt.Println("Namastey")
}
func result(a int, b int) int {
	return a + b
}

// pro functions
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hey love u bro"

}
