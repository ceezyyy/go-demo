package main

import "fmt"

func main() {

	// init
	e1 := &employee{salary: 100}

	e1.giveRaise(0.5)

	fmt.Printf("salary: %f", e1.salary)

}

type employee struct {
	salary float32
}

func (e *employee) giveRaise(percentage float32) {
	e.salary += percentage * e.salary
}
