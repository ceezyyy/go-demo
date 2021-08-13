package main

import "fmt"

func main() {
	res := returnEmptyInterfaceSlice()
	acceptEmptyInterfaceSlice(res)  // [1 2]
	acceptEmptyInterfaceSlice2(res)
}

func returnEmptyInterfaceSlice() []interface{} {
	res := make([]interface{}, 0, 50)
	res = append(res, 1, 2)
	return res
}

func acceptEmptyInterfaceSlice(input []interface{}) {
	fmt.Println(input)
}

func acceptEmptyInterfaceSlice2(input ...interface{}) {

}
