package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var contacts []Contact

type Contact struct {
	Name  string
	Phone string
	Email string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addContact(reader)
		case 2:
			listContacts()
		case 3:
			searchContact(reader)
		case 4:
			deleteContact(reader)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func addContact(reader *bufio.Reader) {
	fmt.Println("enter a name")
	name, _ := reader.ReadString('\n')
	fmt.Println("enter the number")
	phone, _ := reader.ReadString('\n')
	fmt.Println("enter the email")
	email, _ := reader.ReadString('\n')

	contacts = append(contacts, Contact{
		Name:  strings.TrimSpace(name),
		Phone: strings.TrimSpace(phone),
		Email: strings.TrimSpace(email),
	})
	fmt.Println("contact added!")
}
func listContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts available.")
		return
	}
	for i, c := range contacts {
		fmt.Printf("%d. %s - %s - %s\n", i+1, c.Name, c.Phone, c.Email)
	}
}
func searchContact(reader *bufio.Reader) {
	fmt.Print("Enter name to search: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	found := false

	for _, c := range contacts {
		if strings.EqualFold(c.Name, name) {
			fmt.Printf("Found: %s - %s - %s\n", c.Name, c.Phone, c.Email)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("No contact found.")
	}
}
func deleteContact(reader *bufio.Reader) {
	fmt.Print("Enter name to delete: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	for i, c := range contacts {
		if strings.EqualFold(c.Name, name) {
			contacts = append(contacts[:i], contacts[i+1:]...)
			fmt.Println("Contact deleted.")
			return
		}
	}
	fmt.Println("Contact not found.")
}
