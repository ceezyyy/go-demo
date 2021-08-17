package main

import (
	"fmt"
	"time"
)

// Counter -> Squares -> Printer
func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			time.Sleep(1 * time.Second)
			naturals <- x
		}
	}()

	// squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}

}
