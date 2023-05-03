package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	fmt.Println(colors)

	fmt.Println(colors["white"])

	delete(colors, "white")

	fmt.Println(colors)

}
