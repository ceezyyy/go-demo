package main

import (
	"fmt"
	"reflect"
)

func main() {

	// A variable is an addressable storage location that contains a value
	// and it's value may be updated through that address

	x := 2                   // non-addressable, only a copy of integer 2
	a := reflect.ValueOf(2)  // non-addressable
	b := reflect.ValueOf(x)  // non-addressable, a copy of pointer &x
	c := reflect.ValueOf(&x) // non-addressable
	d := c.Elem()            // addressable, derived from c by de-referencing the pointer within it

	// Calling reflect.ValueOf(&x).Elem() to obtain an addressable Value for any variable x

	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())

	// Recover the variable from an addressable reflect.Value
	x = 2
	addr := reflect.ValueOf(&x).Elem()
	xPtr := addr.Addr().Interface().(*int)
	*xPtr = 88
	fmt.Println(x)

	addr.Set(reflect.ValueOf(100.00)) // panic, float is not assignable to int
	fmt.Println(x)

}
