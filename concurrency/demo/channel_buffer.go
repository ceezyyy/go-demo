package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
		fmt.Println("received")
	}()

	fmt.Println("sending")
	ch <- 88
	fmt.Println("sent")

	time.Sleep(2 * time.Second)

}
