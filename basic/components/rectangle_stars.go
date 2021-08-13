package main

import "fmt"

func main() {

	// width = 20, height = 10
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
