package main

import "fmt"

func main() {
	para1 := 1
	para2 := 2
	res := 0
	// 参数 3 传递的是地址
	para3 := &res
	multiply(para1, para2, para3)

	fmt.Printf("res: %d", res)

}

func multiply(a, b int, res *int) {
	*res = a * b
}
