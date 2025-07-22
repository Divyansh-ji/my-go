package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Lets understand a bit of switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open it ")
	case 2:
		fmt.Println("Dice value is 2 and you can't open it ")
	case 3:
		fmt.Println("Dice value is 3 and you can't open it ")
	case 4:
		fmt.Println("Dice value is 4 and you can't open it ")
	case 5:
		fmt.Println("Dice value is 5 and you can,t open it ")
	case 6:
		fmt.Println("Dice value is 6 and you can,t open it")

	default:
		fmt.Println("what was that!")

	}

	// multiple statmwnt in same case

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("yeah it's a weekend")

	default:
		fmt.Println("fuck its a weekday")
	}

}
