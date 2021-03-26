package main

import (
	"fmt"
)

func main() {

	i := foo()
	yr, s := bar()

	fmt.Println(i)
	fmt.Println("year", yr, "book", s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big Brother"
}
