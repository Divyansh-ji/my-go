package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://pkg.go.dev/net/http#Request/coursename=Tiwariji"

func main() {
	fmt.Println("Welcome to under url in go")

	result, _ := url.Parse(myurl)

	//fmt.Println(result.Scheme)
	//	fmt.Println(result.Fragment)
	//fmt.Println(result.Path)
	//fmt.Println(result.Port())
	//fmt.Println(result.Host)

	qparams := result.Query()

	fmt.Println(qparams["coursename"])

	partsurl := &url.URL{
		Scheme:  "http",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user = Divyansh ",
	}
	anotherurl := partsurl.String()
	fmt.Println("hmm", anotherurl)

}
