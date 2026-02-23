package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	age         int
	contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}
	fmt.Println(p1)

	var p2 person
	fmt.Println(p2)

	p2.firstName = "Jane"
	p2.lastName = "Smith"
	p2.age = 25

	fmt.Println(p2)

	jim := person{
		firstName: "Alice",
		lastName:  "Johnson",
		age:       28,
		contactInfo: contactInfo{
			email:   "alice.johnson@example.com",
			zipCode: 12345,
		},
	}
	jim.updateName("Alice Updated")
	jim.print()
	jim.updateNamePointer("Alice Updated Again")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNamePointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
