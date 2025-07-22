package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs in go lang")

	Divyansh := User{"Divyansh", "Divyansh@gmail.com", true, 16}
	fmt.Println(Divyansh)

	// print with structs

	fmt.Printf("Divyansh details are : %+v\n", Divyansh)

	fmt.Printf("Name is %v and email is %v\n", Divyansh.Name, Divyansh.Email)

	person := struct {
		name string
		age  int
	}{

		"Alice",
		30,
	}
	fmt.Println(person.name) // Output: Alice

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
