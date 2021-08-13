package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

}

func sum(vals ...int) (total int) {
	// vals 是 []int
	for _, val := range vals {
		total += val
	}
	return
}
