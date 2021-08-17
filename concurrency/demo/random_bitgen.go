package main

import "fmt"

// unending sequence of 0's & 1's
func main() {

	fmt.Println("main start")
	ch := make(chan int)

	// producer
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	// consumer
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("main end")

}
