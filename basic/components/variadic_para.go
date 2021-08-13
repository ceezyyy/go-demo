package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	testVariadic(sl...)
}

//testVariadic 可变参数
func testVariadic(nums ...int) {
	fmt.Printf("%T", nums)
}
