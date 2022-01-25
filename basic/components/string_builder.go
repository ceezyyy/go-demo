package main

import (
	"fmt"
	"strings"
)

func main() {

	sb := strings.Builder{}

	for i := 0; i < 10000; i++ {
		sb.WriteString("a")
	}

	fmt.Println(sb.String())

}
