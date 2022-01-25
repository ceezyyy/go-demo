package main

import "fmt"

// gopl: ch4
func main() {

	var x []int

	for i := 0; i < 30; i++ {
		x = appendInt(x, i)
		fmt.Printf("cap(x): %d, len(x): %d, x: %v\n", cap(x), len(x), x)
	}

}

func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// reallocating
		zcap := zlen
		if zcap < 2*len(x) { // len(x) 有可能为 0
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z

}
