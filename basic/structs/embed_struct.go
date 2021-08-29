package main

import "fmt"

func main() {

	// init
	b := new(B)
	b.b1 = 1
	b.b1 = 2
	b.a1 = 3
	b.a2 = 4

	fmt.Println(b)

}

// diyStruct 更高级别抽象
type A struct {
	a1, a2 int
}

// B embedded struct
// 体现继承关系
type B struct {
	b1, b2 int
	A
}
