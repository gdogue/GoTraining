package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 20; i++ {
		j := i % 4
		switch {
		case j == 0:
			fmt.Printf("%v is evenly divided by 4\n", i)
		case j == 1:
			fmt.Printf("%v divided by 4 has a remainder of 1\n", i)
		case j == 2:
			fmt.Printf("%v divided by 4 has a remainder of 2\n", i)
		case j == 3:
			fmt.Printf("%v divided by 4 has a remainder of 3\n", i)
		}
	}
}
