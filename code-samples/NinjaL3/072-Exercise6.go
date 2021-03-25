package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%v is evenly divided by 4\n", i)
		}
	}
}
