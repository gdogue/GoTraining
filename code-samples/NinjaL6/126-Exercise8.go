package main

import (
	"fmt"
)

func main() {
	f1 := foo()

	fmt.Println("The value is:", f1())
}

func foo() func() int {
	return func() int {
		return 1984
	}
}
