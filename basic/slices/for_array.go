package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	size := 3
	res := make([][]int, 0)

	length := len(arr)
	for offset := 0; offset < length; offset += size {
		end := minInt(offset+size, length)
		batch := arr[offset:end]
		res = append(res, batch)
	}

	fmt.Println(res)
	fmt.Println(753971603 % 100)
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
