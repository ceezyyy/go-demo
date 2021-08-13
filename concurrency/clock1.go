package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

// gopl: ch8
func main() {

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// use one goroutine per connection
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
