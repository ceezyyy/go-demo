package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("starting the client")

	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Printf("client failed: %v\n", err)
		return
	}

	// get client name
	fmt.Println("what's your name?")
	name, _ := reader.ReadString('\n')
	name = strings.Trim(name, "\n")

	for {

		fmt.Println("Please send you msg: (press q to quit)")
		// read from input
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		// quit
		if input == "q" || input == "Q" {
			return
		}

		// write data to the connection
		msg := []byte(name + " says: " + input)
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Printf("send msg failed: %v\n", err)
		}

	}

}
