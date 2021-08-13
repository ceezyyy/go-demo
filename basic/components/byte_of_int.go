package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := -90
	b := 90

	// 8 bytes
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

}
