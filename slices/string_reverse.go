package main

import "fmt"

func main() {

	s1 := "Google"
	c1 := []byte(s1)

	reverse(c1)

	fmt.Printf("after: %v\n", string(c1))

}

func reverse(str []byte) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
