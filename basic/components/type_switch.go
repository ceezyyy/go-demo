package main

import "fmt"

func main() {
	var a obj = true
	switch v := a.(type) {
	case int:
		fmt.Printf("int: %v, %T", v, v)
	case string:
		fmt.Printf("string: %v, %T", v, v)
	case bool:
		fmt.Printf("bool %v, %T", v, v)
	default:
		fmt.Printf("Unknown type")
	}
}

type obj interface {
}
