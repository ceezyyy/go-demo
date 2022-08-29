package main

import (
	"fmt"
	"go-demo/common-mistakes/init-func/rds"
)

func init() {
	fmt.Println("init main 1")
}

func init() {
	fmt.Println("init main 2")
}

// init() may lead to some issues:
// 1) Limiting error management
// 2) If the init  requires to set a state, it has to be done through global variables (what do you mean?)
func main() {
	_ = rds.Store("foo", "bar")
	return
}
