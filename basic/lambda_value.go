package main

import "fmt"

func main() {
	fv := func() { fmt.Println("hello world") }
	fv()
	fmt.Printf("type: %T\n", fv)
}
