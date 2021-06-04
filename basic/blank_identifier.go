package main

import "fmt"

func main() {
	res1, _, res3 := ThreeValues()
	fmt.Printf("res1: %d, res3: %.2f\n", res1, res3)
}

func ThreeValues() (int, int, float64) {
	return 1, 2, 3.0
}
