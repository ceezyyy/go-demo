package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))

	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))

}
