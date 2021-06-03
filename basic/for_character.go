package main

import "fmt"

func main() {

	// 25 times
	for i := 0; i < 25; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

}
