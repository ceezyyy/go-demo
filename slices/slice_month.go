package main

import "fmt"

func main() {

	// array of month
	sl := [...]string{1: "Jan"}

	for i, v := range sl {
		fmt.Printf("%v: %v\n", i, v)
	}

	fmt.Println(2 * -1)

}
