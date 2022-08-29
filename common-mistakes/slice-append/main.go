package main

import "fmt"

// CAUTION!
// Adding an element using append checks whether the slice is full (len == cap)
// If not full, the append func will add the element by updating the backing array, and incr len by one
func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
