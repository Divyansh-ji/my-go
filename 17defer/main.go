package main

import "fmt"

func main() {
	fmt.Println("hey welcome to go defer ")
	// defer basically skip that line and runs it at the end of the program or at the last braces

	defer fmt.Println("This is the defered line")
	fmt.Println("i like salman")

	//What will happen if there are multiple defer

	defer fmt.Println("Hey i am the 1 defer")
	defer fmt.Println("Hey i am the 2 defer")
	defer fmt.Println("Hey i am the 3 defer")

	//They will follow LIFO !

	myDefer()
	//defer function

}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}

}
