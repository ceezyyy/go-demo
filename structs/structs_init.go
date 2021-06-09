package main

import "fmt"

type interval struct {
	start, end int
}

func main() {

	inter1 := interval{0, 3}
	inter2 := interval{end: 23, start: 6}
	inter3 := interval{end: 88}

	fmt.Println("value:")
	fmt.Printf("%v\n", inter1)
	fmt.Printf("%v\n", inter2)
	fmt.Printf("%v\n", inter3)

	// value 类型, 不同内存地址
	fmt.Println("memory address:")
	fmt.Printf("%p\n", &inter1)
	fmt.Printf("%p\n", &inter2)
	fmt.Printf("%p\n", &inter3)

	inter4 := &interval{start: 2, end: 55}
	inter5 := interval{start: 2, end: 55}
	fmt.Printf("%v\n", inter4)
	fmt.Printf("%v\n", inter5)
	fmt.Printf("%p\n", inter4)
	fmt.Printf("%p\n", &inter5)
}
