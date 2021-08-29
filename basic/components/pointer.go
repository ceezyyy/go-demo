package main

import "fmt"

func main() {

	x := 1
	// p points to x
	p := &x
	fmt.Println(*p) // 1

	*p = 2
	fmt.Println(x) // 2

	var a, b int
	p1 := &a    // addr, not nil
	p2 := &b    // addr, not nil
	var p3 *int // the zero-value of pointer is nil

	fmt.Println(p1 == p2)  // false
	fmt.Println(p1 == nil) // false
	fmt.Println(p2 == nil) // false
	fmt.Println(p3 == nil) // true

}
