package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}

	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	//Create X different GO routines
	for i := 0; i < x; i++ {
		go func() {
			//In each Go routine add Y elements to the channel
			//The elements are integers from 0 to Y-1
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("In gen, ROUTINES", runtime.NumGoroutine())
	}

	return c
}
