package main

import "fmt"

func main() {

	// pipe that connect concurrent goroutines
	msg := make(chan string)

	// By default, channels are unbuffered
	// meaning that they will only accept sends (chan <-), if there is a corresponding receive (<- chan) ready to receive the sent value
	// here: deadlock
	received := <-msg

	// new goroutine
	go func() {
		msg <- "ping"
	}()

	// By default
	// sends and receives block until both the sender and receiver are ready
	// received := <-msg, ok
	fmt.Println(received)

}
