package main

import "fmt"

func main() {
	t := &T{a: 7, b: 2.35, c: "abc\tdef"}
	fmt.Printf("%v", t)
}

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	s := fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
	return s
}
