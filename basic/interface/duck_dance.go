package main

import "fmt"

func main() {
	bird := new(Bird)
	DuckFunc(bird)
}

type IDuck interface {
	Quacking()
	Walking()
}

func DuckFunc(duck IDuck) {
	duck.Quacking()
	duck.Walking()
}

type Bird struct {
}

func (b *Bird) Quacking() {
	fmt.Println("Bird is quacking")
}

func (b *Bird) Walking() {
	fmt.Println("Bird is walking")
}
