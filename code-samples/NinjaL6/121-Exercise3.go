package main

import (
	"fmt"
)

func main() {

	defer snafu()

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)

	fmt.Println("Sum:", sum)

	barsum := bar(xi)

	fmt.Println("barsum:", barsum)
}

func snafu() {
	fmt.Println("I'm done!")
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
