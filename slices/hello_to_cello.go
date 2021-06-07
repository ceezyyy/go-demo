package main

import "fmt"

func main() {

	// string 是不可变的
	s1 := "hello"
	c1 := []byte(s1)
	fmt.Printf("c1: %v\n", c1)

	c1[0] = 'c'
	fmt.Printf("c1: %v\n", c1)

	s2 := string(c1)
	fmt.Printf("s2: %v", s2)

}
