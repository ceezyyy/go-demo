package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// strings are immutable arrays of bytes
	s1 := "a b c"
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

	s2 := "世界"
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))

}
