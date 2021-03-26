package main

import (
	"fmt"
)

func main() {
	f1 := func() {
		fmt.Println("Anonymous func ran")
	}

	f2 := func(x int) {
		fmt.Println("The meaning of life:", x)
	}

	f1()
	f2(42)
}
