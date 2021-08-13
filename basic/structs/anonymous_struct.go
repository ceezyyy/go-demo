package main

import "fmt"

func main() {
	ms := &ms{1, 2, "3"}
	fmt.Println(ms)
}

type ms struct {
	a float32
	int
	string
}
