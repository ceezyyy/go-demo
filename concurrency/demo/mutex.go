package main

import (
	"fmt"
	"sync"
)

// Mutex is a way to guard critical sections of your program
func main() {
	var (
		count int
		lock  sync.Mutex
		wg    sync.WaitGroup
	)

	incr := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("incr: %d\n", count)
	}

	decr := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decr: %d\n", count)
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			incr()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			decr()
		}()
	}
	wg.Wait()

	fmt.Printf("Completed, count: %d\n", count)
}
