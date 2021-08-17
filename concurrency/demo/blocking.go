package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go receiver(ch)
	ch <- 88

	time.Sleep(1 * time.Second)

}

//func sender(out chan int) {
//	out <- 88
//}

func receiver(in chan int) {
	fmt.Println(<-in)
}
