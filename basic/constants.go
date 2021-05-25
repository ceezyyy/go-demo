package main

import "fmt"

// package-level
const s = "constant"

func main() {

	fmt.Println(s)

	const n = 5000000

	const d = 3e20 / n
	fmt.Println(d)



}
