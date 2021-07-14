package main

import (
	"fmt"
)

func main() {

	bool1 := true

	if bool1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	s1 := make([]int, 3)

	s2 := []int{1, 2}

	s3 := append(s1, s2...)

	fmt.Println(s3)

}
