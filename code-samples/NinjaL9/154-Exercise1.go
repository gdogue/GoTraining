package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		fmt.Println("Hello from anonymous function")
		wg.Done()
	}()

	go foo()

	fmt.Println("Hello, playground")
	wg.Wait()
}

func foo() {
	fmt.Println("Hello from foo")
	wg.Done()
}
