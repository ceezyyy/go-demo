package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	f := func(num int) int {
		return num * 10
	}
	res := mapFunc(arr, f)
	fmt.Printf("res: %v", res)
}

func mapFunc(arr []int, f func(num int) int) []int {
	for i := range arr {
		// map
		arr[i] = f(arr[i])
	}
	return arr
}
