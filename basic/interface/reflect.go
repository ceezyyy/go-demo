package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	// reflect.TypeOf() returns an interface value's dynamic type
	// It always return a concrete type
	t := reflect.TypeOf(3)
	fmt.Println(t)

	var w io.Writer = os.Stdout
	t = reflect.TypeOf(w)
	fmt.Println(t)

	fmt.Printf("%T\n", false)
	fmt.Printf("%T\n", 1) // %T: type
	fmt.Printf("%T\n", 1.00)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v) // %v: value
	fmt.Println(v.String())

	v = reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)

}
