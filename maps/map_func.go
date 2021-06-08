package main

import "fmt"

// val 的类型可以是函数
func main() {
	mf := map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
	}
	fmt.Printf("type: %T\n", mf)
	// int -> 函数的内存地址
	fmt.Printf("value of mf: %v", mf)
}
