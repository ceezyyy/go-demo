package main

import (
	"fmt"
	"reflect"
)

func main() {

	// struct
	var secret interface{} = NonKnownType{s1: "LBJ", s2: "AD"}

	fmt.Printf("type: %T\n", secret)
	typ := reflect.TypeOf(secret)
	val := reflect.ValueOf(secret)
	fmt.Printf("type: %v, value: %v\n", typ, val)

	// fields
	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("field: %v\n", val.Field(i))
	}
}

type NonKnownType struct {
	s1, s2 string
}
