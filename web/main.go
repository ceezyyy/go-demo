package main

import (
	"fmt"
	"net/http"
)

// Server
func main() {

	// controller
	http.HandleFunc("/hello", sayHello)

	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Oooops! err: %v\n", err)
		return
	}

}

//sayHello 请求, 响应
func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>Hello World</h1>")
}
