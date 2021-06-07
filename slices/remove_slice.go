package main

import "fmt"

var start = 6
var end = 6

func main() {

	sl1 := []string{"M", "N", "O", "P", "Q", "R"}
	fmt.Printf("origin: %v\n", sl1)

	res := RemoveStringSlice(sl1, start, end)
	fmt.Printf("%v", res)

}

//RemoveStringSlice removes items from index start
// to end in a slice (to: 不包含)
func RemoveStringSlice(src []string, start, end int) []string {
	// corner case
	if start < 0 || end > len(src) || end-start <= 0 {
		return src
	}
	copy(src[start:], src[end:])
	cut := end - start
	return src[:len(src)-cut]
}
