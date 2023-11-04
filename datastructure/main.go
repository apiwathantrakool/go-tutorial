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
	person01 := person{firstName: "Alex", lastName: "Anderson", contact: contactInfo{
		email:   "alex@test.com",
		zipCode: 1234,
	}}
	person01.updateName("Hello")
	person01.print()
}

func (p *person) updateName(newData string) { // *pointer = Give me the value this memory address is pointing at.
	p.firstName = newData
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
