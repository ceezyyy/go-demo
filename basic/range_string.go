package main

import "fmt"

func main() {

	s := "test"
	//non-ascii
	//s := "你好"

	// English char: 1 byte
	// Chinese char: 3 bytes
	for i, c := range s {
		fmt.Printf("Index: %d, value: %c\n", i, c)
	}

}
