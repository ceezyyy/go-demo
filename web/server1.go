package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// global
var listener net.Listener
var err error

func main() {

	var conn net.Conn

	fmt.Println("starting server")
	listener, err = net.Listen("tcp", "localhost: 10000")
	if err != nil {
		fmt.Printf("start listener failed, %v\n", err)
		return
	}

	// infinite loop: 一直接收请求
	for {
		// wait for the connection
		conn, err = listener.Accept()
		if err != nil {
			fmt.Printf("listen req failed, %v\n", err)
			return
		}
		go response(conn)
	}

}

// response from server
func response(conn net.Conn) {

	// infinite loop: read from console
	for {

		buf := make([]byte, 521)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Printf("reading buf failed, %v\n", err)
			return
		}

		input := string(buf)
		switch {
		case strings.Contains(input, "SH"):
			fmt.Println("shutdown server")
			//_ = conn.Close()
			//_ = listener.Close()
			os.Exit(0)
		default:
			fmt.Printf("received data: %v\n", input)
		}
	}

}
