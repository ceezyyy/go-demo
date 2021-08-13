package main

import "fmt"

func main() {

	num := 88

	// 多条件分支判断
	switch {
	case num < 0:
		fmt.Println("negative")
	case num > 0 && num < 10:
		fmt.Println("single digit")
	case num > 10:
		fmt.Println("multi digit")
	}

}
