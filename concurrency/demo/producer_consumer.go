package main

import "fmt"

func main() {

	ch := make(chan int)
	done := make(chan bool)

	go producer(ch)
	go consumer(ch, done)

	<-done
	fmt.Println("done")

}

func producer(ch chan<- int) {
	// 和 consumer 的 for range 搭配使用
	for i := 0; i <= 90; i = i + 10 {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, done chan bool) {
	// 和 producer 的 close() 搭配使用
	for i := range ch {
		fmt.Println(i)
	}
	done <- true
}
