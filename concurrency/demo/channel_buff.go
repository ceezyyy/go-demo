package main

func main() {

	messages := make(chan string, 2)

	// buffer channel
	// sender 不必阻塞, 即使没有对应的 receiver
	messages <- "test"
	messages <- "test"

	//fmt.Println(<-messages)
	//fmt.Println(<-messages)

}
