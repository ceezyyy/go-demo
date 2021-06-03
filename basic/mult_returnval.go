package main

import "fmt"

func main() {
	sum, product, diff := myFunction(1, 2)
	fmt.Print(sum, product, diff)
}

func myFunction(a, b int) (sum, product, diff int) {
	sum = a + b
	product = a * b
	diff = a - b
	return
}
