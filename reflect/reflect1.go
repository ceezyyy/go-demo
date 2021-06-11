package main

import (
	"fmt"
	"reflect"
)

func main() {

	var element float64 = 66.88

	t := reflect.TypeOf(element)
	v := reflect.ValueOf(element)

	fmt.Printf("type: %v\n", t)
	fmt.Printf("value: %v\n", v)
	//fmt.Println(t.Method(0))

}
