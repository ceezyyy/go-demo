package main

import "fmt"

// fallthrough 的使用
func main() {

	k := 6

	switch k {
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
		fallthrough
	case 7:
		fmt.Println("seven")
		fallthrough
	default:
		fmt.Println("others")
	}

}
