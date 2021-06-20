package main

import (
	"fmt"
	"time"
)

// select lets you wait on multiple channel operations
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	//for i := 0; i < 2; i++ {
	// select 中, 只要有一个 receiver ready 就可以
	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
		//}
	}

}
