package main

import (
	"fmt"
)

// var name type = expression
func main() {

	var n int64
	var f float32
	var b bool
	var s string
	var sl []int
	var m map[string]int
	var ch chan int
	var fu func()

	fmt.Println(n)         // 0
	fmt.Println(f)         //0
	fmt.Println(b)         // false
	fmt.Println(s)         // ""
	fmt.Println(sl == nil) // true
	fmt.Println(m == nil)  // true
	fmt.Println(ch == nil) // true
	fmt.Println(fu == nil) // true

}
