package main

import (
	"fmt"
	"math"
)

var (
	// In Go, an integer overflow that can be detected at compile time will generate a compilation error
	//counter int32 = math.MaxInt32 + 1
	counter int32 = math.MaxInt32
)

func main() {
	// At runtime, an integer overflow and underflow are silent, it will not lead to an application panic
	counter++
	fmt.Println(counter)
}
