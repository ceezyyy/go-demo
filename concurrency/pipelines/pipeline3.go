package main

import (
	"fmt"
	"time"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)

}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
