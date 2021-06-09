package main

import "fmt"

type myStruct struct {
	i1 int
	f1 float32
	s  string
}

func main() {

	// init
	ms := new(myStruct)
	ms.i1 = 1
	ms.f1 = 2
	ms.s = "hello struct"

	fmt.Printf("myStruct: %v\n", ms)
	fmt.Println(ms)

}
