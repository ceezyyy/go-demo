package main

import "fmt"

func main() {

	var arr1 [5]int
	var slice1 = arr1[1:3]

	// init
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// slice1
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	fmt.Printf("len(arr1): %d\n", len(arr1))
	fmt.Printf("len(slice1): %d\n", len(slice1))
	fmt.Printf("cap(slice1): %d\n", cap(slice1))

	sl := make([]int, 2)
	fmt.Println(sl)

	for i := 0; i < len(sl); i++ {
		if i == len(sl)-1 { // 最后一个
			sl = append(sl[:i], sl[i+1:]...)
		}
	}

	fmt.Println(sl)
}
