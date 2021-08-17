package main

import "fmt"

func main() {

	ch := make(chan int)
	go fibonacciChan(20, ch)
	// consumer
	for c := range ch {
		fmt.Printf("%v ", c)
	}

}

// producer
func fibonacciChan(term int, ch chan int) {
	for i := 0; i < term; i++ {
		ch <- fibonacci(i)
	}
	close(ch)
}

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
