package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go lang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "peach"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("fruit list is ", len(fruitList))

	//Another way of using arrays

	var vegList = [5]string{"potato", "beans", "panner"}
	fmt.Println("vegy list is :", len(vegList))

}
