package main

import (
	"fmt"
)

func main() {

	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 59, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
