package main

import "fmt"

// range 本身就 receive 了!!!
func main() {

	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"
	close(queue)

	for e := range queue {
		fmt.Println(e)
	}

}
