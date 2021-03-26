package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "My age is", p.age)
}

func main() {

	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p)
	p.speak()
}
