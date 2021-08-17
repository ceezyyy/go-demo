package main

import "fmt"

func main() {

	fmt.Println("main start")
	ch := make(chan int)

	go tel(ch)

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("main done")

}

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}
