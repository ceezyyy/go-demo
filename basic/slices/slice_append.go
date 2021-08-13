package main

import (
	"fmt"
)

// 重要!!!
// slice append
func main() {

	s1 := []int{10, 20, 30, 40, 50}
	s2 := s1[1:3]

	// 扩容时, 当底层数组
	// 1. 有额外容量可用: 修改底层数组
	// 2. 没有额外容量可用: 创建新的底层数组, 并将旧的赋值给新数组
	s2 = append(s2, 60)
	fmt.Println("print s2")
	for _, v := range s2 {
		fmt.Println(v)
	}

	fmt.Println("print s1")
	for _, v := range s1 {
		fmt.Println(v)
	}

	s3 := []int{1, 2, 3, 4}
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	s4 := append(s3, 5)
	fmt.Println(len(s4))
	// 扩容为原来的 2 倍
	fmt.Println(cap(s4))

	s5 := []string{"a", "b", "c", "d"}
	s6 := append(s5, "e")
	fmt.Println(cap(s6))

	// 设置长度和容量相同
	s7 := s5[1:3:3]
	fmt.Println(s7)
	s7 = append(s7, "e")
	if cap(s7) == 4 {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}

	s8 := []int{1, 2}
	s9 := []int{3, 4, 5, 6, 7, 8, 9}
	s10 := append(s8, s9...)
	fmt.Println(s10)
	fmt.Println(len(s10))
	fmt.Println(cap(s10))
}
