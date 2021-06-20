package main

import (
	"fmt"
	"time"
)

// do sth in the future
func main() {

	timer1 := time.NewTimer(2 * time.Second)

	// channel of timer
	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	// 手动停止
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

}
