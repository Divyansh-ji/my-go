package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs in go lang")

	Divyansh := User{"Divyansh", "Divyansh@gmail.com", true, 16}
	fmt.Println(Divyansh)

	// print with structs

	fmt.Printf("Divyansh details are : %+v\n", Divyansh)

	fmt.Printf("Name is %v and email is %v\n", Divyansh.Name, Divyansh.Email)
	Divyansh.getStatus()
	Divyansh.NewMail()
	r := rect{width: 10, height: 5}
	fmt.Println("so answer is", r)
	fmt.Println("so calculated area is", r.area())

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

func (u User) getStatus() {
	fmt.Println("Is user active", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("the new email", u.Email)
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}
