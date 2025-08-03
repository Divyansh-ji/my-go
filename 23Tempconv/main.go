package main

import "fmt"

func main() {
	var choice int
	var temp float64

	fmt.Println("hey welcome to my temperature converter")
	for {
		fmt.Println("\n===Temperatur==")

		fmt.Println("press 1 to convert c to f")
		fmt.Println("press 2 to convert f to c")
		fmt.Println("press 3 to exist")
		fmt.Println("enter the choices")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("enter the temperature in celcicus")
			fmt.Scan(&temp)
			fmt.Printf("%.2f °C = %.2f °F\n", temp, ctof(temp))
		case 2:
			fmt.Println("Enter the temperature to convert in farehenite")
			fmt.Scan(&temp)
			fmt.Printf("%.2f °F = %.2f °C\n", temp, ftoc(temp))
		case 3:
			fmt.Println("extsing....")
			return
		default:
			fmt.Println("Invalid choice")

		}

	}
}

func ctof(c float64) float64 {
	return (c * 9 / 5) + 32

}
func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
