package main

import (
	"fmt"
	"runtime"
)

// todo
func main() {
	msg := receiveMsg()
	_ = getMsgType(msg)
}

func receiveMsg() []byte {
	return make([]byte, 1000000)
}

func getMsgType(msg []byte) []byte {
	tmp := msg[:5]
	printAlloc()
	return tmp
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
