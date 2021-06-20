package main

import (
	"fmt"
	"time"
)

func main() {

	// buffered
	c1 := make(chan string, 1)

	go func() {
		// 2 秒的操作
		time.Sleep(2 * time.Second)
		c1 <- "msg one"
	}()

	// Select proceeds with the first receive that’s ready
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout msg one")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "msg two"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout msg two")
	}

}
