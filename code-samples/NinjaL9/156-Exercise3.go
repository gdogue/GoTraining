package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	var wg sync.WaitGroup

	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		v := counter
		runtime.Gosched()
		v++
		counter = v
		wg.Done()
	}

	fmt.Println("Before Wait...")
	wg.Wait()
	fmt.Println("Done, Counter:", counter)
}
