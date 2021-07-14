package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()
	router.GET("/", Index)

	fmt.Println("starting server")

	log.Fatal(http.ListenAndServe(":8888", router))

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello world")
}
