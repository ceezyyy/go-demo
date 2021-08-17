package main

import "fmt"

// Ref: Go maps in action https://blog.golang.org/maps
func main() {

	// A nil map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	fmt.Println(len(m1) == 0)

	//m1["sd"] = 1 // runtime panic

	m := make(map[string]int)

	// if the requested key doesn't exist
	// we get the value type's zero value
	fmt.Println(m["sdsd"])

	// "v, ok" tests for the existence of a key
	v, ok := m["testKey"]
	fmt.Println(v)
	fmt.Println(ok)

}
