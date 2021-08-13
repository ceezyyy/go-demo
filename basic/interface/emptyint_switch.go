package main

import "fmt"

type testStr string

var str testStr

func main() {
	str = "test"
	testType(str)
	testType(1)
	testType("sss")
}

// 入参为 empty interface
func testType(any interface{}) {
	switch t := any.(type) {
	case int:
		fmt.Printf("type: %T, val: %v\n", t, t)
	case float32:
		fmt.Printf("type: %T, val: %v\n", t, t)
	case string:
		fmt.Printf("type: %T, val: %v\n", t, t)
	case bool:
		fmt.Printf("type: %T, val: %v\n", t, t)
	default:
		fmt.Println("Oooops")
		fmt.Printf("type: %T, val: %v\n", t, t)
	}
}
