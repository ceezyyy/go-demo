package main

import "fmt"

func main() {
	fmt.Println(testFunc1())

}

func testFunc1() (res []int) {
	res = append(res, 1, 2)
	return
}
