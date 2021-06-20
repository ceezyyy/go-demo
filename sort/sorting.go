package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"a", "b", "c"}
	fmt.Printf("%p\n", s)

	// in-place
	sort.Strings(s)

	fmt.Println(s)
	fmt.Printf("%p\n", s)

	ints := []int{4, 5, 6, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	isSorted := sort.IntsAreSorted(ints)
	fmt.Println(isSorted)

}
