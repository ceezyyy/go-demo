package main

import "fmt"

func main() {
	printrec(10)
}

func printrec(i int) {
	// from 10 to 1
	if i < 1 {
		return
	}
	fmt.Println(i)
	printrec(i - 1)
}
