package main

import "fmt"

// close a channel 表示没有更多的 value 将会被发送
func main() {

	jobs := make(chan int, 5)
	// sync
	done := make(chan bool)

	// receiver
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received", j)
			} else {
				fmt.Println("all received")
				// 记得 sync
				done <- true
				// 记得 return
				return
			}
		}
	}()

	// sender (main goroutine)
	for i := 0; i < 2; i++ {
		jobs <- i
		fmt.Println("sent", i)
	}
	close(jobs)
	fmt.Println("all jobs sent")

	// syn, 程序结束
	<-done

}
