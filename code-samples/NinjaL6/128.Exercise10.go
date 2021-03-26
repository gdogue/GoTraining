package main

import (
	"fmt"
)

func main() {

	x := 42
	fmt.Println(x)
	f := func() {
		x := 1984
		fmt.Println(x)
	}
	f()
	fmt.Println(x)
}
