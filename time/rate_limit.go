package main

import (
	"fmt"
	"time"
)

// 控制速率
// 1) 提高资源利用率
// 2) 保持服务稳定
func main() {

	request := make(chan int, 5)

	// sender
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)

	// send the time on the channel after each tick
	limiter := time.Tick(time.Second)

	for r := range request {
		// 控制接收的速率
		<-limiter
		fmt.Println("request: ", r, time.Now())
	}

}
