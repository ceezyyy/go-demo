package main

import "fmt"

func main() {

	var arr [16]int

	for i := 0; i <= 15; i++ {
		arr[i] = i
	}

	for i := range arr {
		fmt.Println(arr[i])
	}

}
