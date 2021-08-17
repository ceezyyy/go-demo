package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// Receiver
	go func() {
		time.Sleep(15 * time.Second)
		x := <-ch
		fmt.Println("receiving", x)
	}()
	fmt.Println("sending")
	ch <- 10
	fmt.Println("sent")
}
