package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var x int = 1
	//var d time.Duration = 1 * time.Second

	fmt.Println(Any(x))

}

func Any(v interface{}) string {
	return formatAtom(reflect.ValueOf(v))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // array, struct, interface
		return v.Type().String() + "value"
	}
}
