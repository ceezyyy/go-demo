package main

import "fmt"

func main() {

	// init
	outer := new(outer)
	outer.c = 3
	outer.d = 4
	outer.int = 5
	outer.a = 1
	outer.b = 2

	fmt.Println(outer)

}

type inner struct {
	a, b int
}

type outer struct {
	c, d  int
	int   // anonymous field
	inner // anonymous field
}
