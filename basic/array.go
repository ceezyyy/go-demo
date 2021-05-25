package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)

	a[2] = 3
	fmt.Println(a)

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

}
