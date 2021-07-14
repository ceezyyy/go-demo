package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("starting the server")

	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("listener failed: ", err.Error())
		return
	}

	// infinite for-loop
	for {
		// 接收请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept req failed: ", err.Error())
			return
		}
		go doServerThing(conn)
	}

}

func doServerThing(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		// read data from connection
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("reading buf failed: ", err.Error())
			return
		}
		fmt.Printf("received data: %v\n", string(buf))
	}
}
