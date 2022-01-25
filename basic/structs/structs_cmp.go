package main

import (
	"fmt"
	"reflect"
)

type diyStruct struct {
	field1 int
	field2 string
}

func main() {

	a1 := &diyStruct{
		field1: 1,
		field2: "a",
	}

	a2 := &diyStruct{
		field1: 1,
		field2: "a",
	}

	fmt.Println(a2)
	fmt.Println(a2.field1)
	fmt.Println(a2.field2)

	fmt.Println(a1 == a2) // false

	// Two values of identical type are deeply equal
	fmt.Println(reflect.DeepEqual(a1, a2)) // true

}
