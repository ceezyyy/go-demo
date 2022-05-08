package main

import (
	"fmt"
	"unsafe"
)

const (
	YearStart  = 6
	MonthStart = 10
	DayStart   = 12
	DayEnd     = 14
)

func main() {

	a := -90
	b := 90

	// 8 bytes
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

}

type Test struct {
	*CC

	A int
}

type CC struct {
	cc int
}
