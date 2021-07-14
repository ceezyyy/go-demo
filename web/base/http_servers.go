package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8888", nil)
}

// handler
// response writer is used to fill in the HTTP response
func hello(w http.ResponseWriter, req *http.Request) {

	_ = req.ParseForm()
	id := req.Form.Get("id")
	fmt.Printf("type: %T\n", id)

	idInt, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(idInt)
	fmt.Printf("type: %T", idInt)

	// 输出到 response writer
	fmt.Fprint(w, id)
}
