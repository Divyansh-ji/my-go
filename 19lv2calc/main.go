package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a float64, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("CLI Calculator (type 'exit' to quit)")

	for {
		fmt.Print("\nEnter calculation (e.g., 5 + 3): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit condition
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator...")
			break
		}

		// Split input into parts
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid input! Format: number operator number (e.g., 2 + 3)")
			continue
		}

		// Convert strings to numbers
		a, err1 := strconv.ParseFloat(parts[0], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers! Please try again.")
			continue
		}

		op := parts[1]
		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result:", result)
	}
}
