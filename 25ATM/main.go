package main

import "fmt"

type Account struct {
	ID      int
	PIN     int
	Balance int
}

var account []Account
var Id int = 1

func main() {
	var choice int

	for {
		fmt.Println("\n=== ATM Simulator ===")
		fmt.Println("1. Create Account")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			CreateAccount()
		case 2:
			Login()
		case 3:
			fmt.Println("👋 Exiting ATM...")
			return
		default:
			fmt.Println("❌ Invalid choice")
		}
	}

}
func CreateAccount() {

	var pin int
	var balance int
	fmt.Println("ENTER THE PIN")
	fmt.Scan(&pin)
	fmt.Println("Enter the balance")
	fmt.Scan(&balance)

	account = append(account, Account{ID: Id, PIN: pin, Balance: balance})
	fmt.Println("Your account has been created ", Id)
	Id++

}
func accountMenu(acc *Account) {
	var choice int

	for {
		fmt.Println("\n--- Account Menu ---")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Logout")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("💰 Your Balance: %d\n", acc.Balance)
		case 2:
			deposit(acc)
		case 3:
			withdraw(acc)
		case 4:
			fmt.Println("👋 Logged out.")
			return
		default:
			fmt.Println("❌ Invalid choice")
		}
	}
}

func deposit(acc *Account) {
	var amount int
	fmt.Println("Enter the amount to deposite")
	fmt.Scan(&amount)
	acc.Balance += (amount)
	fmt.Println("New bankbalance is", acc.Balance)

}
func withdraw(acc *Account) {
	var amount int
	fmt.Println("enter the amount to be withdraw")
	fmt.Scan(&amount)
	if amount > acc.Balance {
		fmt.Println("❌ Insufficient funds.")
	} else {
		acc.Balance -= amount
		fmt.Println("✅ Withdrawal successful. New Balance:", acc.Balance)
	}

}
func Login() {
	var id int
	var pin int

	fmt.Println("enter the id")
	fmt.Scan(&id)
	fmt.Println("enter the pin")
	fmt.Scan(&pin)
	for i, _ := range account {
		if account[i].ID == id && account[i].PIN == pin {
			fmt.Println("Login is succesfull ✅")
			accountMenu(&account[i])
			return
		} else {
			fmt.Println("Invalid pin or account id")
		}

	}
}
