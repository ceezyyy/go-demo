package main

import "fmt"

func main() {

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1

	// 打印虚拟地址
	fmt.Printf("arr1: %p, arr2: %p\n", &arr1, &arr2)

	// 修改 arr2
	for i := range arr2 {
		arr2[i] += 1
	}

	// arr1
	for i := range arr1 {
		fmt.Printf("%d ", arr1[i])
	}

}
