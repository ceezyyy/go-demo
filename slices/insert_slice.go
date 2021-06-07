package main

import "fmt"

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"g", "o", "l", "a", "n", "g"}
	for i := 0; i <= len(s2); i++ {
		res := InsertStringSlice(s1, s2, i)
		fmt.Printf("%v\n", res)
	}
}

//InsertStringSlice inserts a slice into another slice
// at a certain index
func InsertStringSlice(intersection, dest []string, index int) []string {

	// 1. 扩容
	res := make([]string, len(intersection)+len(dest))

	// 2. copy [start, index, border, end]
	copy(res, dest[:index])
	border := index + len(intersection)
	copy(res[index:border], intersection)
	copy(res[border:], dest[index:])

	// 3. 赋值
	dest = res
	return dest

}
