package main

import "fmt"

func main() {

	// Copy returns the number of elements copied
	s1From := []int{1, 2, 3}
	s2To := make([]int, 10)

	copy(s1From, s2To)

	fmt.Printf("s1From: %v\n", s1From)
	fmt.Printf("s2To: %v\n", s2To)
	fmt.Printf("Address of s1From: %p\n", s1From)
	fmt.Printf("Address of s2To: %p\n", s2To)

	// 不够就分配新的数组 -> 内存地址不一样
	s3 := []int{1, 2, 3}
	fmt.Printf("Address of s3: %p\n", s3)
	s3 = append(s3, 4)
	fmt.Printf("Address of s3: %p\n", s3)

	// 容量足够则直接插入 -> 内存地址不变
	s4 := make([]int, 3, 6)
	fmt.Printf("Address of s4: %p\n", s4)
	s4 = append(s4, 1, 2)
	fmt.Printf("Address of s4: %p\n", s4)

}
