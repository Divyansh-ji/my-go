package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//we will learn about customeunmarshall
//UnmarshalJSON() lets you customize how JSON becomes your struct
// understand by example
//Let’s say we receive a name in uppercase, but we want to store it lowercase in the struct.

type User struct {
	Name string
}

func (u *User) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name string
	}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	u.Name = strings.ToLower(temp.Name)

	return nil

}
func main() {
	jsonData := []byte(`{"name": "Divyansh "}`)

	var u User
	err := json.Unmarshal(jsonData, &u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(u)

}
