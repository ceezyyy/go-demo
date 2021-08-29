package main

import "fmt"

// gopl: ch4
func main() {
	a := []int{1, 2, 3, 4}
	reverse(a)
	fmt.Println(a)
}

func reverse(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}
