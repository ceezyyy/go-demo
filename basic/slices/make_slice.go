package main

import "fmt"

func main() {

	slice1 := make([]int, 5)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 2 * i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%v ", slice1[i])
	}
	fmt.Println()

	fmt.Printf("len: %d, cap: %d\n", len(slice1), cap(slice1))

	slice2 := make([]byte, 5)
	fmt.Printf("len: %d, cap: %d\n", len(slice2), cap(slice2))

	slice3 := slice2[2:4]
	fmt.Printf("len: %d, cap: %d\n", len(slice3), cap(slice3))

	slice4 := []byte{'p', 'o', 'e', 'm'}
	slice5 := slice4[2:]
	fmt.Printf("slice5: %v\n", slice5)
	slice4[0] = 'a'
	slice4[2] = 'b'
	// same memory address
	fmt.Printf("slice4: %v\n", slice4)
	fmt.Printf("slice5: %v\n", slice5)

}
