package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int64

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Done, Counter:", counter)
}
