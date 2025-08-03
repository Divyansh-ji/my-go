package main

import "fmt"

func main() {
	var choice int
	var amount int

	for {
		fmt.Println("\n====CURRCONV===")
		fmt.Println("Enter 1 for Inrtousd")
		fmt.Println("Enter 2 for Usdtoinr")
		fmt.Println("press 3 for exist")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Please enter the amount")
			fmt.Scan(&amount)
			fmt.Print(InrtoUsd(float64(amount)))
		case 2:
			fmt.Println("please enter the amout ")
			fmt.Scan(&amount)
			fmt.Println(UsdtoInr(float64(amount)))
		case 3:
			fmt.Println("existing")
			return
		default:
			fmt.Println("Inavlid")

		}
	}

}
func InrtoUsd(inr float64) float64 {
	return inr * 0.012
}
func UsdtoInr(usd float64) float64 {
	return 87 * usd
}
