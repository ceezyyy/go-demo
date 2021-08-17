package main

import (
	"fmt"
	"time"
)

// gopl: ch8
// You needn't close every channel when you're finished it
// It's only necessary to close a channel when it is important
// to tell the receiving goroutines that all data have been sent
func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
			time.Sleep(500 * time.Millisecond)
		}
		//
		close(naturals)
	}()

	// squarer
	go func() {
		for n := range naturals {
			squares <- n * n
		}
		close(squares)
	}()

	// printer (in main goroutine)
	for s := range squares {
		fmt.Println(s)
	}

}
