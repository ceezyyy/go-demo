package main

import "fmt"

// 重要!!!
func main() {

	num := 101

	switch num {
	// case 之后不用 {}
	case 98, 99: // 多个值匹配
		fmt.Println("It's equal to 98 or 99")
	case 100:
		fmt.Println("It's equal to 100")
		//default:
		//	fmt.Println("It's not equal to 98, 99 or 100")
	}

}
