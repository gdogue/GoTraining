package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {

	p1 := person{
		first:   "Gary",
		last:    "Richardson",
		flavors: []string{"strawberry", "vanilla"},
	}

	p2 := person{
		first:   "Ann",
		last:    "Richardson",
		flavors: []string{"strawberry", "Chocolate"},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.flavors {
		fmt.Println("\t", v)
	}

	fmt.Println(p2.first, p2.last)
	for _, v := range p2.flavors {
		fmt.Println("\t", v)
	}

}
