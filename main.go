package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred!"`)
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}

	sa := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	saySomething(sa)
}
