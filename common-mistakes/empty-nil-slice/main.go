package main

import "fmt"

var i = 0

// First, one of the main diff between nil and an empty slice is regarding allocations
// Second, regardless if a slice is nil, the append built-in func does work
func main() {
	var s []string
	log(s)

	s = []string(nil) // not most widely used
	log(s)

	s = []string{}
	log(s)

	s = make([]string, 0)
	log(s)
}

// empty: len(sli) == 0
// nil: sli == nil
func log(s []string) {
	i++
	fmt.Printf("%d, isEmpty: %t, isNil: %t\n", i, len(s) == 0, s == nil)
}
