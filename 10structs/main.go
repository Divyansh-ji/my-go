package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs in go lang")

	Divyansh := User{"Divyansh", "Divyansh@gmail.com", true, 16}
	fmt.Println(Divyansh)

	// print with structs

	fmt.Printf("Divyansh details are : %+v\n", Divyansh)

	fmt.Printf("Name is %v and email is %v", Divyansh.Name, Divyansh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
