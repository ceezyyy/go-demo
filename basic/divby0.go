package main

import "fmt"

func main() {
	a := 9.00
	res := a / 0.0
	fmt.Println(res) // +Inf
}
