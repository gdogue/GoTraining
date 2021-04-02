package main

import (
	"fmt"
)

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is a custom error: %v\n", ce.err)
}

func main() {
	c1 := customErr{
		err: "Custom Error 1001"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
