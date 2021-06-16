package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	p1 := point{1, 2}
	p2 := &point{1, 2}

	fmt.Printf("%v\n", p1)
	// 输出 fields
	fmt.Printf("%+v\n", p1)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)

	// string
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")

	// pointer
	fmt.Printf("%p\n", p2)

	// 只返回 string
	s := fmt.Sprintf("a %s", "str")
	fmt.Println(s)

}
