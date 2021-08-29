package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("x: %d\n", x/x)  // panic if x == 0
	defer fmt.Printf("%d\n", x) // stack
	f(x - 1)
}
