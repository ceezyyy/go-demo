package main

import (
	"fmt"
)

// 使用 select-default 实现 non-blocking send/receiver
func main() {

	messages := make(chan string)
	signal := make(chan string)

	select {
	// blocking, 走 default 分支
	case msg := <-messages:
		fmt.Println("received", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello world"
	select {
	// blocking, 走 default 分支
	case messages <- msg:
		fmt.Println("sent")
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received", msg)
	case sig := <-signal:
		fmt.Println("received", sig)
	default:
		fmt.Println("no activity")
	}

}
