package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main start")

	go longSleep()
	go shortSleep()

	time.Sleep(10 * time.Second)
	fmt.Println("main end")

}

func longSleep() {
	fmt.Println("long sleep start")
	time.Sleep(5 * time.Second)
	fmt.Println("long sleep end")
}

func shortSleep() {
	fmt.Println("short sleep start")
	time.Sleep(2 * time.Second)
	fmt.Println("short sleep end")
}
