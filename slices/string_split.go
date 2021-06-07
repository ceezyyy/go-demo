package main

import "fmt"

func main() {
	s1 := "golang"
	// test
	for i := 0; i < len(s1); i++ {
		sub1, sub2 := Split(s1, i)
		fmt.Printf("sub1: %s, sub2: %s\n", sub1, sub2)
	}
}

func Split(str string, pos int) (sub1, sub2 string) {
	return str[:pos], str[pos:]
}
