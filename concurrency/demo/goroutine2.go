package main

import (
	"fmt"
	"time"
)

func main() {

	//time.Sleep(2 * time.Second)

	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	// 让 sendData() & getData() 能完成
	// 放前面的话, sleep 后直接结束执行
	time.Sleep(1 * time.Second)

}

//sendData 发送者
func sendData(ch chan string) {
	ch <- "Beijing"
	ch <- "Washington"
	ch <- "Paris"
	ch <- "Tokyo"
}

//getData 接收者
func getData(ch chan string) {
	var s string
	// 死循环
	for {
		s = <-ch
		fmt.Println(s)
	}
}
