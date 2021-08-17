package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Printf("create listener failed|err: %v\n", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("get connection failed|err: %v\n", err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
	_ = conn.Close()
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	_, _ = fmt.Fprintln(conn, strings.ToUpper(shout))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(conn, shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(conn, strings.ToLower(shout))
}
