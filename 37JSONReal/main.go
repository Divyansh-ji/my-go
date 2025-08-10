package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Change struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Timestamp int64  `json:"timestamp"`
	User      string `json:"user"`
}

func main() {

	req, err := http.Get("https://stream.wikimedia.org/v2/stream/recentchange")

	if err != nil {
		fmt.Println("Error has been occured ", err)
	}
	defer req.Body.Close()
	scanner := bufio.NewScanner(req.Body) //convert chunks into line

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data") {
			data := strings.TrimPrefix(line, "data")
			var c Change

			err := json.Unmarshal([]byte(data), &c)
			if err != nil {
				fmt.Println("Error has been occured", err)
			}
			fmt.Println("the readed json live is", c.ID, c.Title, c.Timestamp)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)

	}
}
