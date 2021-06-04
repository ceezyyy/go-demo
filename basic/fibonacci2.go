package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		val, pos := fibonacci2(i)
		fmt.Printf("value: %d, position: %d\n", val, pos)
	}
}

func fibonacci2(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci2(n - 1)
		v2, _ := fibonacci2(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}
