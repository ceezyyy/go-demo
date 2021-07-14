package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	// init router
	// 注册 handler
	router := httprouter.New()
	router.POST("/body", showBody)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe("localhost:8888", router))

}

func showBody(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// 获得 body 长度
	size := r.ContentLength

	// Body 实现了 io.Reader 和 io.Closer
	// 需要用 []byte 接收
	body := make([]byte, size)

	r.Body.Read(body)

	fmt.Fprintf(w, string(body))

}
