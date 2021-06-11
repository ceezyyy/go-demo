package main

import "fmt"

func main() {

	car1 := &Car{wheelCount: 88}
	car2 := &Car{wheelCount: 10}

	fmt.Printf("%v, %v", car1.wheelCount, car2.wheelCount)

}

type Car struct {
	// private
	wheelCount int
}

func (car *Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	// 继承 Car
	Car
}
