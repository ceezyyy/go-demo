package main

import (
	"fmt"
	"sync"
	"time"
)

// 重要!!!
// wait group: 等待所有 goroutines 完成
func main() {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		// counter++
		wg.Add(1)
		go worker2(i, &wg)
	}

	// blocked until counter == 0
	wg.Wait()

}

func worker2(id int, wg *sync.WaitGroup) {

	// counter--
	defer wg.Done()

	fmt.Printf("worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)

}
