package main

import "fmt"

func main() {

	fmt.Println("welcome to american express")
	Divyansh := Account{"Divyansh", 21, 250000}
	fmt.Println(Divyansh)

	//use of withdraw
	Divyansh.Withdraw(50000)

	//use of deposit
	Divyansh.Deposite(50000)

	//we can make for any use

	Anushka := Account{"Anushka ", 23, 200000}
	fmt.Println(Anushka)
	Anushka.Withdraw(9800)
}
func (a *Account) Deposite(amount int) {
	if amount > 0 {
		a.Balance += amount
		fmt.Println("The New balannce is", a.Balance)

	} else {
		fmt.Println("Deposite amount must be postive")

	}
}
func (b *Account) Withdraw(amount int) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive.")
		return
	}
	if b.Balance > amount {
		b.Balance = b.Balance - amount
		fmt.Println("The remaining balance is", b.Balance)
	} else {
		fmt.Println("Insufficent balance")
	}
}

type Account struct {
	Owner   string
	Age     int
	Balance int
}
