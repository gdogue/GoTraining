package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func updateAge(pAddr *person, newAge int) {
	(*pAddr).age = newAge //Or can just use "pAddr.age = newAge"
}

func main() {

	p := person{
		first: "Ann",
		last:  "Richardson",
		age:   51,
	}
	fmt.Println(p.age)
	updateAge(&p, p.age+1)
	fmt.Println(p.age)
}
