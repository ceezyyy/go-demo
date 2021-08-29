package main

import "fmt"

// The example demonstrates the mechanism of panic and defer
// Ref: https://go.dev/blog/defer-panic-and-recover Defer, Panic, and Recover
func main() {
	f()
	fmt.Println("return normally from f")
}

func f() {

	// 1. recover 只在 defer 函数有效
	// 2. 只针对当前 goroutine!!
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f", r)
		}
	}()

	fmt.Println("Calling g")
	g(0)
	fmt.Println("return normally from g") // 程序无法走到此行, 已 panic

}

func g(i int) {

	if i > 1 {
		panic("panic!!")
	}
	defer fmt.Println("defer in g", i)

	fmt.Println("g here!", i)
	g(i + 1)

}
