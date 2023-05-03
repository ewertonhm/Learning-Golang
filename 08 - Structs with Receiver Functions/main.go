package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p1 := person{
		firstName: "Monty",
		lastName:  "Python",
		contactInfo: contactInfo{
			email:   "mail@go.com.br",
			zipCode: 12345,
		},
	}
	// pointer1 := &p1
	// pointer1.updateName("Mounty")
	p1.updateName("Mounty")
	p1.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	// (*p).firstName = newFirstName
	p.firstName = newFirstName
}
