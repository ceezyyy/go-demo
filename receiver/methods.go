package main

import "fmt"

// Methods is a function that acts on variable of a certain type
func main() {

	// init
	ti1 := &TwoInts{100, 200}
	ti2 := &TwoInts{1, 2}

	res1, res2 := ti1.Add(), ti2.Add()
	fmt.Println(res1)
	fmt.Println(res2)

}

type TwoInts struct {
	a, b int
}

func (ti *TwoInts) Add() (res int) {
	res = ti.a + ti.b
	return
}
