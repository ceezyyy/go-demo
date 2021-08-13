package main

import "fmt"

func main() {

	var arr [5]int

	// 这里的 "i" 指的是 index
	for i := range arr {
		fmt.Println(i)
	}
	//
	//for i := 0; i < len(arr); i++ {
	//	fmt.Printf("index: %d, val: %d\n", i, arr[i])
	//}
	aa := [...]string{"a", "b", "c"}
	for i := range aa {
		fmt.Printf("index: %d, val: %s\n", i, aa[i])
	}

}
