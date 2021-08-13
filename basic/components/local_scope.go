package main

import "fmt"

// global
var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	fmt.Println(a)
}

func m() {
	// global 和 local 变量可以重名
	// local 优先级更高
	a := "0"
	//a = "0"
	fmt.Println(a)
}
