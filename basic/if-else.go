package main

import "fmt"

func main() {

	// parenthesis: not required
	// braces: required
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 can be divided by 4")
	}

	// 可以定义块局部变量 / 初始化函数
	// num 为块局部变量
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multi digits")
	}

	if a, b := 1, 2; a < b {
		fmt.Println("a", a)
	} else {
		fmt.Println("b", b)
	}

}
