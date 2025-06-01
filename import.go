package main

import (
	"fmt"
	net "net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library")

	resp, err := net.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status: ", resp.Status)
}
