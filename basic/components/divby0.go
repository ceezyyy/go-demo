package main

import (
	"fmt"
	"time"
)

func main() {
	a := 9.00
	res := a / 0.0
	fmt.Println(res) // +Inf

	sl := make([]interface{}, 0, 10)
	for i := 0; i < 2; i++ {
		sl = append(sl, 1, 2)
	}

	today := time.Now().Format("2006-01-02")
	fmt.Println(today)

	fmt.Println(time.Now().Unix())

	var arr []int
	arr = append(arr, 1)
	fmt.Println(arr)


}
