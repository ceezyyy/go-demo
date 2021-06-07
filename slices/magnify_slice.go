package main

import "fmt"

func main() {

	sl1 := []int{1, 2, 3}
	fmt.Printf("len(sl1): %d\n", len(sl1))

	sl1 = enlarge(sl1, 5)
	fmt.Printf("len(sl1): %d\n", len(sl1))
	fmt.Printf("sl1: %v\n", sl1)

}

//enlarge len(s) * factor
func enlarge(s []int, factor int) (newS []int) {
	newS = make([]int, len(s)*factor)
	copy(newS, s)
	return
}
