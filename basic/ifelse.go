package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)

	str := "abc"
	str = ""
	if str == "" || len(str) == 0 {
		fmt.Println("empty string")
	} else {
		fmt.Println("not an empty string")
	}

	fmt.Println(Abs(100))
	fmt.Println(Abs(-100))

	fmt.Println(isGreater(1, 88))
	fmt.Println(isGreater(88, 88))
	fmt.Println(isGreater(89, 88))

}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
