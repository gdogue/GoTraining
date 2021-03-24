package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 11

	z := x == y
	fmt.Println(z)
	z = x <= y
	fmt.Println(z)
	z = x >= y
	fmt.Println(z)
	z = x != y
	fmt.Println(z)
	z = x > y
	fmt.Println(z)
	z = x < y
	fmt.Println(z)
}
