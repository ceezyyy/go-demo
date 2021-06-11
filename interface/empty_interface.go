package main

import (
	"fmt"
)

var i = 100
var s = "test"

type Ms struct {
	a, b float32
}

type Any interface{}

func main() {

	var any Any

	// int
	any = i
	fmt.Printf("val: %v, type: %T\n", any, any)

	// string
	any = s
	fmt.Printf("val: %v, type: %T\n", any, any)

	// struct
	ms := &Ms{a: 1, b: 2}
	any = ms
	fmt.Printf("val: %v, type: %T\n", any, any)

}
