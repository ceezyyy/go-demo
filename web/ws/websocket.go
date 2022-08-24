package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("read msg failed")
				return
			}

			fmt.Printf("%s send: %s, type: %d\n", conn.RemoteAddr(), string(msg), msgType)

			if err = conn.WriteMessage(msgType, msg); err != nil {
				fmt.Println("write msg failed")
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./websockets.html")
	})

	fmt.Println("start successfully")
	_ = http.ListenAndServe(":8080", nil)
}
