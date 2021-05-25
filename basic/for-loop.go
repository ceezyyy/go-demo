package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println("i", i)
		i++
	}

	for j := 10; j < 15; j++ {
		fmt.Println("j", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}


}
