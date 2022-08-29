package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte // pointer field
}

func main() {
	elems := make([]Foo, 1000)
	printAlloc()

	for i := 0; i < len(elems); i++ {
		elems[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc() // todo

	//two := keepFirstTwoElementsOnly(elems)
	two := keepFirstTwoElementsOnlyCopy(elems)

	// If the element is a pointer or a struct with pointer fields,
	// the elements won't be reclaimed by the GC
	runtime.GC()

	printAlloc()

	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(elems []Foo) []Foo {
	return elems[:2]
}

func keepFirstTwoElementsOnlyCopy(elems []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, elems)
	return res
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
