package main

import "fmt"

func main() {

	employee := new(Employee)
	id := employee.Id()
	fmt.Println(id)

	employee.SetId(88)
	id = employee.Id()
	fmt.Println(id)

}

// Base 父类
type Base struct {
	id int64
}

func (b Base) Id() int64 {
	return b.id
}

func (b *Base) SetId(id int64) {
	b.id = id
}

//Person 人
type Person struct {
	Base
	FirstName string
	LastName  string
}

// Employee 员工
type Employee struct {
	Person
	salary float32
}
