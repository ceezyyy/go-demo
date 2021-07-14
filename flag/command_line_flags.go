package main

import (
	"flag"
	"fmt"
)

// command-line flag parsing
func main() {

	// go run command_line_flags.go -word hello-world
	wordPtr := flag.String("word", "foo", "a string")
	intPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	//flag.StringVar()

	// 解析
	flag.Parse()
	fmt.Println(*wordPtr)
	fmt.Println(*intPtr)
	fmt.Println(*boolPtr)

}
