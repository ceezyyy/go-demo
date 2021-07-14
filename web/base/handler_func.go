package main

import (
	"fmt"
	"net/http"
)

func main() {

	// init handler
	http.HandleFunc("/test", test)

	// server
	server := http.Server{Addr: "localhost:8888"}

	// start server
	server.ListenAndServe()

}

// handler function, 注意和原始的写法有何不同?
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test succeeded!")
}
