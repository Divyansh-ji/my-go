package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hey welcome to files in go lang")

	content := "This needs to go in the file"

	// for creation we use os operations
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is ", length)
	defer file.Close()

	//
	readFile("./mylcogofile.txt")
}

// for reading or writing we use ioutil
func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is\n", string(databytes))
}
