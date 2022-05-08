package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		start = time.Now()
		c     = make(chan int)
	)

	go func() {
		time.Sleep(3 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking")
	select {
	case <-c:
		fmt.Printf("Unblocked, %v sec later\n", time.Since(start))
	}

	fmt.Println("------------------")

	c1 := make(chan struct{})
	close(c1)
	c2 := make(chan struct{})
	close(c2)

	// The Go runtime cannot infer your problem space or why you placed a group of channels together in to a select statement
	// Because of this, the best thing the Go runtime can hope to do is to work well in the average case
	var cnt1, cnt2 int
	for i := 0; i < 100; i++ {
		select {
		case <-c1:
			cnt1++
		case <-c2:
			cnt2++
		}
	}

	fmt.Println(cnt1, cnt2)
}
