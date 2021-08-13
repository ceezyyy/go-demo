package main

import "fmt"

func main() {

	s := "old"
	p := &s
	fmt.Printf("s: %s\n", s)
	fmt.Printf("ptr: %p\n", p)

	// Symbol "*" gives the value which the pointer is pointing to
	*p = "new"
	fmt.Printf("s: %s\n", s)
	fmt.Printf("ptr: %p\n", p)
	fmt.Printf("*p: %s\n", *p)

}
