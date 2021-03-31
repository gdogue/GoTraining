package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var myLock sync.Mutex

	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			myLock.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			myLock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Done, Counter:", counter)
}
