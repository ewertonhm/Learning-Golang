package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	p1 := person{
		firstName: "Monty",
		lastName:  "Python",
		contact: contactInfo{
			email:   "mail@go.com.br",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", p1)
}
