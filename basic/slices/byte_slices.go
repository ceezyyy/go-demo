package main

import (
	"fmt"
)

func main() {

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("value: %v, type: %T, addr: %p\n", b, b, b)

	b1 := b[1:4]
	fmt.Printf("value: %v, type: %T, addr: %p\n", b1, b1, b1)

	// pointer: 首地址
	b2 := b[:2]
	fmt.Printf("value: %v, type: %T, addr: %p\n", b2, b2, b2)

	b3 := b[2:]
	fmt.Printf("value: %v, type: %T, addr: %p\n", b3, b3, b3)

	b4 := b[:]
	fmt.Printf("value: %v, type: %T, addr: %p\n", b4, b4, b4)

}
