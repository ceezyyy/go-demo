package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 注册为 handler
	http.HandleFunc("/showHeader", showHeader)

	// server
	server := http.Server{Addr: "localhost:8888"}

	// start server
	fmt.Println("starting server")
	server.ListenAndServe()

}

func showHeader(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintln(w, header)
}
