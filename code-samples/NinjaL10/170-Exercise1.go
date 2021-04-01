package main

import (
	"fmt"
)

func main() {
	//Func Literal approach - commented out
	/* c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	*/

	//Buffered approach
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
