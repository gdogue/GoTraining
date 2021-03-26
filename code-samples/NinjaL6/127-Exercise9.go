package main

import (
	"fmt"
)

func main() {

	math(5, 10, add)
	math(5, 10, multiply)
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func multiply(x int, y int) {
	fmt.Println(x * y)
}

func math(x int, y int, f func(op1 int, op2 int)) {
	f(x, y)
}
