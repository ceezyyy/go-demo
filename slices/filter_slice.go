package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := func(num int) bool {
		if num%2 == 0 {
			return true
		}
		return false
	}
	res := Filter(sl, f)
	fmt.Printf("res: %v", res)
}

func Filter(sl []int, fn func(int) bool) (res []int) {
	// for loop
	for _, v := range sl {
		// filter
		if fn(v) {
			res = append(res, v)
		}
	}
	return
}
