package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))

	fmt.Println(len(dst), cap(dst))

	// CAUTION!
	copy(dst, src)

	fmt.Println(dst)
}
