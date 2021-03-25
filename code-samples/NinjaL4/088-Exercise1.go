package main

import (
	"fmt"
)

func main() {

	x := [5]int{10, 20, 30, 40, 40}

	for i, v := range x {
		fmt.Println("Index:", i, "Value:", v)
	}

	fmt.Printf("%T\n", x)

}
