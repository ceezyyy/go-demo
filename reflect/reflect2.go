package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := 3.14

	// 传 copy
	typCopy := reflect.TypeOf(x)
	valCopy := reflect.TypeOf(x)

	// 传地址, 防止只修改 copy
	typAddr := reflect.TypeOf(&x)
	valAddr := reflect.ValueOf(&x)

	fmt.Println("copy:")
	fmt.Printf("type: %v, value: %v\n", typCopy, valCopy)

	fmt.Println("address:")
	fmt.Printf("type: %v, value: %v\n", typAddr, valAddr)

	fmt.Println("can set or not:")
	//fmt.Printf("copy: %v", valCopy.CanSet())
	fmt.Printf("addr: %v", valAddr.CanSet())

	// 1) the value that the interface v contains
	// 2) the value that pointer v points to
	valAddr = valAddr.Elem()
	fmt.Printf("type: %v, value: %v\n", typAddr, valAddr)
	valAddr.SetFloat(3.1415926535)
	fmt.Printf("type: %v, value: %v\n", typAddr, valAddr)

}
