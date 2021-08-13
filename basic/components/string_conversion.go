package main

import (
	"fmt"
	"strconv"
)

// 函数错误处理
func main() {

	//var s1 = "test"
	var s2 = "11"

	// 初始化 & 赋值
	num, err := strconv.Atoi(s2)
	if err != nil {
		// 错误处理
		fmt.Println("String cannot be converted to integer")
		return
	}
	// Normal case
	fmt.Printf("After converted: %d\n", num)

}
