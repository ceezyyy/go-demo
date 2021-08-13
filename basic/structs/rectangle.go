package main

import "fmt"

func main() {

	r1 := new(Rectangle)
	r1.width = 8
	r1.length = 10

	area, perimeter := Area(r1), Perimeter(r1)
	fmt.Printf("area: %d, perimeter: %d", area, perimeter)

}

type Rectangle struct {
	length, width int
}

func Area(rec *Rectangle) (area int) {
	area = rec.length * rec.width
	return
}

func Perimeter(rec *Rectangle) (perimeter int) {
	perimeter = (rec.length + rec.width) * 2
	return
}
