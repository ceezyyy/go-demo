package main

import "fmt"

func main() {

	// semaphore
	signal := make(chan int)
	fmt.Println("main start")

	go func(a, b int, ch chan int) {
		// get sum
		ch <- a + b
	}(10, 20, signal)

	fmt.Println(<-signal)
	fmt.Println("done")

}
