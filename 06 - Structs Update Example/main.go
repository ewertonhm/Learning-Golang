package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var p1 person
	fmt.Printf("%+v\n", p1)
	p1.firstName = "Monty"
	fmt.Printf("%+v\n", p1)
	p1.lastName = "Python"
	fmt.Printf("%+v\n", p1)
}
