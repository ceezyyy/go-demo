package main

import (
	"fmt"
	"time"
)

var counter int

// Excellent for repeated tasks
func main() {

	fmt.Println("ticker tutorial")

	ticker := time.NewTicker(1 * time.Second)

	// C <-chan Time // The channel on which the ticks are delivered
	for _ = range ticker.C {
		counter++
		fmt.Printf("ticker %d\n", counter)
	}

}
