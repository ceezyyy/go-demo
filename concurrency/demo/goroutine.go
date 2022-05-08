package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memoryConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var (
		c    <-chan interface{}
		wg   sync.WaitGroup
		noop = func() {
			defer wg.Done()
			<-c
		}
		loop = 1e6
	)

	before := memoryConsumed()
	for i := loop; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memoryConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/loop/1000)
}
