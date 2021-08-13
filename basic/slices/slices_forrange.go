package main

import "fmt"

func main() {

	slice1 := make([]int, 4)

	slice1[0] = 0
	slice1[1] = 1
	slice1[2] = 2
	slice1[3] = 3

	for i, v := range slice1 {
		// v 是 copy
		fmt.Printf("%d, %d\n", i, v)
	}

	seasons := []string{"spring", "summer", "autumn", "winter"}

	// index & value
	for i, season := range seasons {
		fmt.Printf("%d, %s\n", i, season)
	}

	// only value
	for _, season := range seasons {
		fmt.Printf("%s\n", season)
	}

	// only index
	for i := range seasons {
		fmt.Printf("%d\n", i)
	}

	// 创建 slice
	items := []int{10, 20, 30, 40, 50}
	for _, item := range items {
		// 只改变 copy
		item *= 2
	}
	fmt.Println(items)

	for i := range items {
		items[i] *= 2
	}
	fmt.Println(items)

}
