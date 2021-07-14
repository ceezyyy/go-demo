package main

import (
	"fmt"
	"time"
)

func main() {

	//ts := time.Now().Unix()
	//fmt.Println(ts)

	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.Format("2006-01-02 15:04:05"))

}
