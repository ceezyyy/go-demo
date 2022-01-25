package main

import (
	"fmt"
	"sync"
)

// Common Mistakes: https://github.com/golang/go/wiki/CommonMistakes
func main() {

	var out []*int
	// In Go, the loop iterator variable is a single variable that takes different values in each loop iteration
	for i := 0; i < 3; i++ {
		//i := i copy i into a new variable
		// We append the same address which eventually contains the last value that was assigned to i
		out = append(out, &i)
	}

	fmt.Println(*out[0], *out[1], *out[2])
	fmt.Println(out[0], out[1], out[2])

	var outArr [][]int
	for _, arr := range [][1]int{{1}, {2}, {3}} {
		outArr = append(outArr, arr[:])
	}

	wg := sync.WaitGroup{}
	values := []int{1, 2, 3, 4, 5}

	for i := range values {
		val := values[i] // copy into a new variable

		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(val)
		}()

	}

	wg.Wait()

	fmt.Println()
	for i := 1; i <= 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

}
