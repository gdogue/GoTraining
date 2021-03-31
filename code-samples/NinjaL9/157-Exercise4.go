package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	var myLock sync.Mutex

	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		myLock.Lock()
		v := counter
		runtime.Gosched()
		v++
		counter = v
		myLock.Unlock()
		wg.Done()
	}

	fmt.Println("Before Wait...")
	wg.Wait()
	fmt.Println("Done, Counter:", counter)
}
