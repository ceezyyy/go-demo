package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var client net.Conn
var reader *bufio.Reader
var input string

func main() {

	var err error

	fmt.Println("start client")
	client, err = net.Dial("tcp", "localhost: 10000")
	if err != nil {
		fmt.Printf("start client failed, %v\n", err)
		return
	}

	reader = bufio.NewReader(os.Stdin)
	// infinite loop: read from console
	for {

		fmt.Printf("Send msg to server: (press q to quit)\n")

		// handle input
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("client send msg failed, %v\n", err)
			return
		}
		input = strings.Trim(input, "\n")
		if input == "q" || input == "Q" {
			return
		}

		// write msg
		_, err = client.Write([]byte(input))
		if err != nil {
			fmt.Printf("write msg failed, %v\n", err)
			return
		}

	}
}
