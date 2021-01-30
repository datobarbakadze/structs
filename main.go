package main

import "fmt"

type contactInfo struct {
	zipCode string
	email   string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson",
		contact: contactInfo{
			zipCode: "055",
			email:   "bdato269@g.com",
		},
	}
	alex.updateName("hitehre")
	alex.print()

}

func (p *person) updateName(newFirstname string) {
	p.firstName = newFirstname
}

func (p person) print() {
	fmt.Println(p)
}
