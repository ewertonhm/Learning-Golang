package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Value of", key, "is", value)
	}
}
