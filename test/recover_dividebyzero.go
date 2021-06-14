package main

import "fmt"

func main() {
	fmt.Println("start")
	test()
	fmt.Println("end")
}

func badCall() {
	a, b := 100, 0
	res := a / b
	fmt.Println(res)
}

// 测试函数
func test() {
	defer func() {
		if e := recover(); e != nil {
			// 异常处理
			fmt.Printf("Error %s in main.badCall() with no para\n", e)
			return
		}
	}()
	badCall()
	fmt.Println("after badCall()")
}
