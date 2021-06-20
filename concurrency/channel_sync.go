package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go worker1(done)

	// receiver 会阻塞, 直到从 channel 获取了消息
	<-done

}

// goroutine
func worker1(done chan bool) {

	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}
