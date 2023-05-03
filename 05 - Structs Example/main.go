package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{
		firstName: "Tom",
		lastName:  "Cat",
	}
	p2 := person{"Jerry", "Mouse"}

	fmt.Println(p1)
	fmt.Println(p2)
}
