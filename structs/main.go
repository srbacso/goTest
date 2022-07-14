package main

import (
	"fmt"
)

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
	var stephen person

	fmt.Println("Starting")

	alex := person{firstName: "Alex", lastName: "Anderson"}

	alex.lastName = "Smith"

	stephen.firstName = "Stephen"
	stephen.lastName = "Bacso"

	stephen2 := person{
		firstName: "Stephen",
		lastName:  "Bacso",
		contact: contactInfo{
			email:   "stephen.bacso@gmail.com",
			zipCode: 90210,
		},
	}

	pSteve := &stephen

	fmt.Println(alex)
	fmt.Printf("%+v\n", stephen)
	fmt.Printf("%+v\n", stephen2)

	pSteve.updateName("Steve")
	pSteve.print()
	fmt.Println("Ending")
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
