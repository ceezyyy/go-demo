package main

import (
	"fmt"
	"net/http"
)

func main() {

	// handler
	handler := &MyHandler{}

	// server, 转发到 MyHandler
	server := http.Server{Addr: "localhost:8888", Handler: handler}

	server.ListenAndServe()

}

type MyHandler struct {
}

// type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")

}
