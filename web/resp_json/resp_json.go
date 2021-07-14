package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	// init router
	router := httprouter.New()
	router.GET("/user", jsonExample)

	fmt.Println("starting server")

	http.ListenAndServe("localhost:8888", router)

}

// handler
func jsonExample(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	user := &User{Name: "LBJ", Email: "awesome.mail", Age: 1}
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("json encoding failed, ", err)
	}

	//log.Fatal(w.Write(data))
	if _, err := w.Write(data); err != nil {
		log.Println("write json failed, ", err)
	}
}

type User struct {
	Name  string
	Email string
	Age   int
}
