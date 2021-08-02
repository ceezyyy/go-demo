package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println(a)
}
