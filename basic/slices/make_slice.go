package main

import "fmt"

type A struct {
	status int
	Data   B
}

type B struct {
	status int
}

func main() {

	var a A
	fmt.Println(a)
	fmt.Println(a.Data)
	fmt.Println(a.Data.status)

}
