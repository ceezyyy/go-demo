package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	weight = []int{2649, 1700, 220, 1600, 1, 50, 2820, 930, 30}
)

func main() {
	results := RandN(weight, 10000)
	gifts := make(map[int]int)

	for _, res := range results {
		gifts[res]++
	}

	fmt.Println(gifts)
}

func RandN(weights []int, n int) []int {
	rand.Seed(time.Now().UnixNano())

	sum := 0
	for _, v := range weights {
		sum += v
	}

	res := make([]int, 0)
	for ; n > 0; n-- {
		r := rand.Intn(sum)
		for i, v := range weights {
			if r >= v {
				r -= v
			} else {
				res = append(res, i)
				break
			}
		}
	}

	return res
}
