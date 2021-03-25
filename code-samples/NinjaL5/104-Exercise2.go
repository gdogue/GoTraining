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

	m := map[string]person{}

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

	m[p1.first] = p1
	m[p2.first] = p2

	for k, v := range m {
		fmt.Println("Key:", k)
		fmt.Println("\t", v.first, v.last)
		for i, v2 := range v.flavors {
			fmt.Println("\t\t", i, v2)
		}

	}

}
