package main

import (
	"fmt"
	"strings"
)

// 创建 structs 的三种方式
func main() {

	// 方式一
	var p1 Person
	p1.lastName = "a"
	p1.firstName = "b"
	upPerson(&p1)

	// 方式二
	p2 := new(Person)
	p2.lastName = "c"
	p2.firstName = "d"
	upPerson(p2)

	p3 := &Person{lastName: "e", firstName: "f"}
	p3.lastName = "e"
	p3.firstName = "f"
	upPerson(p3)

	// res
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p3: %v\n", p3)

}

type Person struct {
	lastName, firstName string
}

//upPerson 入参为地址
func upPerson(p *Person) {
	p.lastName = strings.ToUpper(p.lastName)
	p.firstName = strings.ToUpper(p.firstName)
}
