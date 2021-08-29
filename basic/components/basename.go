package main

import "fmt"

func main() {
	res := basename("/a/b/c.go")
	fmt.Println(res)
}

func basename(s string) string {

	// remove prefix "/"
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// remove suffix "."
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s

}
