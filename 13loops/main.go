package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	day := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(day)

	for index, day := range day {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	ram := 2
	for ram < 10 {
		if ram == 2 {
			goto lco
		}
		fmt.Println("Value is", ram)
		ram++
	}
lco:
	fmt.Println("jumpinf in ssdsfba")

	// most used

	for _, day := range day {
		fmt.Printf("index is and value is %v\n", day)
	}
}
