package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// abort channel
	abort := make(chan int)
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- 1
	}()

	fmt.Println("Commencing countdown, press return to abort")

	// multiplexing
	// wait on multiple channels operations
	select {
	case <-time.After(10 * time.Second):
	// do nothing
	case <-abort:
		fmt.Println("launch aborted")
		return
	}

	launch()

}

func launch() {
	fmt.Println("Take off!!!")
}
