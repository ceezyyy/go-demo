package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fv := func() { fmt.Println("hello world") }
	fv()
	fmt.Printf("type: %T\n", fv)

	rand.Seed(time.Now().UnixNano())
	max := 10
	min := 1
	for i := 0; i < 100; i++ {
		random := rand.Intn(max-min) + min
		fmt.Println(random)

	}
}
