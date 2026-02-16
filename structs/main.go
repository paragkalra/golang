package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
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
}
