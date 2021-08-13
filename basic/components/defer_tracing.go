package main

import "fmt"

func main() {
	B()
}

func trace(s string) {
	fmt.Printf("Entering: %s\n", s)
}

func untrace(s string) {
	fmt.Printf("Leaving: %s\n", s)
}

func A() {
	trace("A")
	defer untrace("A")
	fmt.Println("In func A")
}

func B() {
	trace("B")
	defer untrace("B")
	fmt.Println("In func B")
	A()
}
