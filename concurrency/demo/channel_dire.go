package main

import "fmt"

// 单方向的 channel
func main() {

	pings := make(chan string)
	pongs := make(chan string)

	// 并发两个 goroutine, unbuffered channel 没有问题
	// 若串行执行, unbuffered channel 会导致 deadlock
	go ping(pings, "hello world")

	go pong(pings, pongs)

	fmt.Println(<-pongs)

}

func ping(ping chan<- string, msg string) {
	// receiver
	ping <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	// receiver from "ping"
	// ping: sender
	msg := <-ping
	// send to main goroutine
	pong <- msg
}
