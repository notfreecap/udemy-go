package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	p := Person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact:   ContactInfo{email: "aa@a.com", zipCode: 94000},
	}

	p.updateName("demo")
	p.print()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(firstName string) {
	(*p).firstName = firstName
}
