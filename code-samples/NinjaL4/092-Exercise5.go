package main

import (
	"fmt"
)

func main() {

	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}
