package main

import (
	"fmt"
	"time"
)

const (
	jobNum  = 10
	workers = 10
)

func main() {

	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)

	// 2 个 worker 并发执行
	for i := 0; i <= workers; i++ {
		go worker(i, jobs, results)
	}

	// 下发任务
	for i := 0; i < jobNum; i++ {
		jobs <- i
	}
	// 显式关闭 channel
	close(jobs)

	// 查看结果
	for i := 0; i < jobNum; i++ {
		<-results
	}

}

//worker 并发执行任务
func worker(id int, jobs <-chan int, results chan<- int) {
	// receive and send to "results"
	for job := range jobs {
		time.Sleep(2 * time.Second)
		fmt.Printf("worker %d done job %d\n", id, job)
		results <- job * 2
	}
}
