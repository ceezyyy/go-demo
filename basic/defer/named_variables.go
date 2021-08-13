package main

import "fmt"

func main() {
	fmt.Printf("res: %d", triple(1))
}

func double(x int) (res int) {
	defer func() {
		fmt.Printf("the return value of double: %d\n", res)
	}()
	res = x * 2
	return
}

func triple(x int) (res int) {
	defer func() {
		fmt.Printf("the return value of triple: %d\n", res+x)
	}()
	res = double(x)
	return
}
