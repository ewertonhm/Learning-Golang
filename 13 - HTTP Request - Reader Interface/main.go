package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	body := make([]byte, 51200)
	n, _ := resp.Body.Read(body)
	resp.Body.Close()

	fmt.Println("Response status:\n", resp.Status)
	fmt.Println("Response body:\n", string(body))
	fmt.Println("Response size in bytes:\n", n)

}
