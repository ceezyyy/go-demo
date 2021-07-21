package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	then := time.Date(2021, 7, 20, 9, 0, 60, 651387237, time.UTC)
	fmt.Println(then)

	fmt.Println(then.Weekday())

	// Compare two times
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	diff1 := now.Sub(then)
	diff2 := then.Sub(now)
	// Duration 带正负
	fmt.Println(diff1.Hours())
	fmt.Println(diff2.Hours())

}
