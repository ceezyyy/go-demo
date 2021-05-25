package main

import (
	"fmt"
)

// short variable declaration: declare & init
// var declaration: 1) type 2) assigned later
func main() {

	// Go will infer the type of initialized variables
	var a = "initial"
	fmt.Println(a)
	var b = true
	fmt.Println(b)

	// "zero-value" of the type
	var c int
	fmt.Println(c)
	var d string
	fmt.Println(d) // ""
	var e bool
	fmt.Println(e) // false

	// var declares 1 or more variables
	var f, g, h int
	fmt.Println(f, g, h)
	var i, j, k = true, "f", 0.0
	fmt.Println(i, j, k)

	// short variable declarations
	// name := expression
	l := 1
	fmt.Println(l)

	m, n := 100, 200
	fmt.Println(m, n)

}
