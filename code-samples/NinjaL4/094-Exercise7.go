package main

import (
	"fmt"
)

func main() {

	x1 := []string{"James", "Bond", "Shaken"}
	x2 := []string{"Miss", "Money", "Stirred"}

	x3 := [][]string{x1, x2}
	fmt.Println(x3)

	for i, v := range x3 {
		fmt.Println("Record:", i)
		for j, v2 := range v {
			fmt.Println("\tIndex:", j, "Value:", v2)
		}

	}
}
