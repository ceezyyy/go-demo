package main

import "fmt"

// For more, read https://blog.golang.org/defer-panic-and-recover
func main() {
	fmt.Println(myFunc())
}

func myFunc() int {

	fmt.Println("start")
	res := 0

	defer func() {
		if res == 1 {
			fmt.Println("res = 1")
		}
	}()

	res = 1

	return res
}
