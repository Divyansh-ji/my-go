package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app!")

	// Create a reader
	reader := bufio.NewReader(os.Stdin)

	// Ask user for rating
	fmt.Print("Please rate our pizza between 1 and 5: ")

	// Read input from user
	input, _ := reader.ReadString('\n')

	// Trim and convert to float
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number:", err)
	} else {
		fmt.Println("Thanks for rating:", numRating)
		fmt.Println("1 added to your rating:", numRating+1)
	}
}
