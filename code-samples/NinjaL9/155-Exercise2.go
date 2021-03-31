package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("I am", p.first, p.last, "My age is", p.age)
}

type human interface {
	speak()
}

func saySomething(p human) {
	p.speak()
}

func main() {

	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	saySomething(&p)
	//saySomething(p)  DOES NOT WORK as the speak function is expecting a pointer

}
