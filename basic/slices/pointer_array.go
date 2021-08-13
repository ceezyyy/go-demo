package main

import "fmt"

func main() {
	input := [3]int{1, 2, 3}
	f(input)
	fp(&input)
}

// 传 copy
func f(arr [3]int) {
	fmt.Println(arr)
}

// 传地址
func fp(arr *[3]int) {
	fmt.Println(arr)
}
