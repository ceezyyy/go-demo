package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(filter("a", "b", "c"))
	fmt.Println(113419298 / 16 % 64)
	fmt.Println(int(24 * time.Hour.Seconds()))
}

func filter(elements ...string) bool {
	all := strings.Join(elements, ",")
	return strings.Contains(all, "d")
}
