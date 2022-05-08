package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		count int
		once  sync.Once
		wg    sync.WaitGroup
	)

	incr := func() {
		count++
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			once.Do(incr)
		}()
	}
	wg.Wait()
	fmt.Println(count)

	test()
}

func test() {
	var (
		count int
		once  sync.Once
	)

	incr := func() {
		count++
	}
	decr := func() {
		count--
	}

	// if once.Do(f) is called multiple times, only the first call will invoke f,
	// even if f has a different value in each invocation.
	once.Do(decr)
	once.Do(incr)
	fmt.Println(count)
}
