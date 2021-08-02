package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//source := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(source)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Intn(100))
	}

}
