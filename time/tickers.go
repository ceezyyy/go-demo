package main

import (
	"fmt"
	"time"
)

// do sth repeatedly
func main() {

	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				//default:
				//	fmt.Println("default here")
			}
		}
	}()

	time.Sleep(10 * time.Second)
	//ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

}
