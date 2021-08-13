package main

import "fmt"

func main() {
	slice1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	min, max := minSlice(slice1), maxSlice(slice1)
	fmt.Printf("min: %d, max: %d", min, max)
}

func minSlice(arr []int) (minimum int) {
	minimum = arr[0]
	for i := range arr {
		if arr[i] < minimum {
			minimum = arr[i]
		}
	}
	return
}

func maxSlice(arr []int) (maximum int) {
	maximum = arr[0]
	for i := range arr {
		if arr[i] > maximum {
			maximum = arr[i]
		}
	}
	return
}
