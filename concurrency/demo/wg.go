package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg          sync.WaitGroup
		salutations = []string{"hello", "greetings", "good day"}
	)
	for _, salutation := range salutations {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// the loop exits before any goroutines begin running,
			// so salutation is transferred to the heap holding a reference to the last value in our string slice
			fmt.Println(salutation)
		}()
	}
	wg.Wait()

	fmt.Println("--------------")

	for _, salutation := range salutations {
		wg.Add(1)
		go func(salutation string) { // pass a copy of "salutation"
			defer wg.Done()

			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
