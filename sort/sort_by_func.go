package main

import (
	"fmt"
	"sort"
)

func main() {

	fruits := byLength{"aaa", "bb", "c"}

	fmt.Printf("origin addr: %p\n", fruits)

	sort.Sort(fruits)
	fmt.Println(fruits)

	fmt.Printf("modified addr: %p\n", fruits)

}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	// 改写逻辑
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
