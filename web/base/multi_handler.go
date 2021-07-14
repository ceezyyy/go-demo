package main

import (
	"fmt"
	"net/http"
)

func main() {

	// init handler
	hello := &helloHandler{}
	world := &worldHandler{}

	// server
	server := http.Server{Addr: "localhost:8888"}

	// router
	http.Handle("/hello", hello)
	http.Handle("/world", world)

	// start server
	server.ListenAndServe()

}

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type worldHandler struct {
}

func (h worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}
