package main

import "fmt"

// 起到 clean-up 的作用
func main() {

	fmt.Println("top")
	// 程序结束时执行
	defer myFunc()
	fmt.Println("bottom")

}

func myFunc() {
	fmt.Println("defer")
}
