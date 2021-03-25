package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%v is evenly divided by 4\n", i)
		} else if i%4 == 2 {
			fmt.Printf("%v divided by 4 has a remainder of 2\n", i)
		} else {
			fmt.Printf("%v is not evenly divided by 4\n", i)
		}
	}
}
