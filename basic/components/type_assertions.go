package main

import "fmt"

// interface{} -> 具体类型
func main() {

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// 需判断 ok, 否则无法转换会 panic!!
	f, ok := i.(float32)
	if ok {
		fmt.Println(f)
	}

}
