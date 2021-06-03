package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	res, err := MySqrt(-100)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("res: %.2f", res)
}

func MySqrt(input float64) (res float64, err error) {
	if input < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	res = math.Sqrt(input)
	return
}
