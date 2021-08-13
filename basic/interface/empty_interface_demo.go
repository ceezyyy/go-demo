package main

import "fmt"

func main() {

	var i interface{}
	fmt.Printf("%v, %T\n", i, i)

	i = 42
	fmt.Printf("%v, %T\n", i, i)

	i = "hello"
	fmt.Printf("%v, %T\n", i, i)

}
