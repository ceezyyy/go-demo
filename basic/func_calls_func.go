package main

import "fmt"

var s string

func main() {
	s = "G"
	fmt.Println(s)
	f1()
}

func f1() {
	s = "0"
	fmt.Println(s)
	f2()
}

func f2() {
	fmt.Println(s)
}
