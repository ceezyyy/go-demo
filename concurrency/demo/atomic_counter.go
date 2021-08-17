package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops int32

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				atomic.AddInt32(&ops, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(ops)

}
