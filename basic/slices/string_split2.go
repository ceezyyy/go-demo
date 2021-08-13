package main

import "fmt"

func main() {
	s := "google"
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	fmt.Printf("res: %v", s2+s1)
}
