package main

import "fmt"

func main() {

	// init
	m1 := make(map[string]int)
	m1["one"] = 1
	m1["two"] = 2

	// test if key exists
	if val, isPresent := m1["two"]; isPresent {
		fmt.Printf("val: %v", val)
	}

}
