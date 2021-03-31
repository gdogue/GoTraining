package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	var counter int64
	wg.Add(100)
	for i := 0; i < 100; i++ {

		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
		wg.Done()
	}

	fmt.Println("Before Wait...")
	wg.Wait()
	fmt.Println("Done, Counter:", counter)
}
