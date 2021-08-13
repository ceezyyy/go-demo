package main

import "fmt"

func main() {
	s := new(Simple)
	s.Set(100)
	fmt.Println(s.Get())
}

type Simpler interface {
	Get() int
	Set(n int)
}

type Simple struct {
}

func (s *Simple) Get() int {
	return 88
}

func (s *Simple) Set(n int) {
	fmt.Println("Set() here!")
}
