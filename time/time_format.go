package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t.Format("2006-01-02"))

	fmt.Println(int64(24 * time.Hour.Seconds()))

}
