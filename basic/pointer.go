package main

import "fmt"

func main() {

	var i int = 1
	// intP points to variable "i"
	var intP *int = &i
	fmt.Printf("Integer: %d, its address: %p\n", i, intP)

	// nil: 声明但未赋值
	var testNil *int
	fmt.Println(testNil)

	// var == *(&var)

}
