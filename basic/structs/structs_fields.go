package main

import "fmt"

type myStruct struct {
	i1 int
	f1 float32
	s  string
}

func main() {

	var ms myStruct

	// zero-value
	fmt.Println(ms.i1)
	fmt.Println(ms.f1)
	fmt.Println(ms.s)

}
