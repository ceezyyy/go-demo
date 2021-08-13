package main

import (
	"fmt"
)

// if-else 初始化可见性
func main() {

	var first = -9

	if first < 0 {
		fmt.Println("negative")
	} else if first > 0 && first < 10 {
		fmt.Println("single digit")
	} else {
		fmt.Println("multi dight")
	}

	var second int
	if second = 9; second < 0 {
		fmt.Println("Negative")
	} else if second > 0 && second < 10 {
		fmt.Println("Single digit")
	} else {
		fmt.Println("Multi digit")
	}

	// 使用 ":=" 对变量进行 for 初始化时
	// 该变量只在 if-else {} 中可见
	if third := 99; third < 0 {
		fmt.Println("Negative")
	} else if third > 0 && third < 10 {
		fmt.Println("Single digit")
	} else {
		fmt.Println("Multi digit")
	}

}
