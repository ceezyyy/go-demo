package main

import (
	"fmt"
	"strings"
)

// person.go 变式
func main() {

	// init
	p1 := new(Person2)
	p1.lastName = "a"
	p1.firstName = "b"

	// 传的是 copy
	upPerson2(*p1)
	// p1 不变
	fmt.Printf("p1: %v", p1)

}

type Person2 struct {
	lastName, firstName string
}

//upPerson2 入参类型为 value type
func upPerson2(p Person2) {
	p.lastName = strings.ToUpper(p.lastName)
	p.firstName = strings.ToUpper(p.firstName)
}
