package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	*InnerMsg
	Name  string
	Body  string
	Time  int64
	Extra int
}

type InnerMsg struct {
	Inner string
}

// 内存 -> 可写入磁盘/通过网络传输到别的机器上
// Ref: JSON and Go https://blog.golang.org/json
func main() {

	// Encoding
	msg := Message{
		Name: "LBJ",
		Body: "Hello world",
		Time: time.Now().UnixNano(),
	}
	b, _ := json.Marshal(msg)
	fmt.Printf("b: %v\n", b)
	// []uint8/[]byte 与 string 互转
	fmt.Printf("type of b: %T\n", b)
	fmt.Println(string(b))

	// Decoding
	c := []byte(`{"Name":"AD", "Body": "AD here!", "Time":1294706395881547000, "Inner": "inner msg here"}`)
	var m Message
	_ = json.Unmarshal(c, &m)
	// print fields
	fmt.Println(m.Name)
	fmt.Println(m.Body)
	fmt.Println(m.Time)
	fmt.Println(m.Extra)
	fmt.Println(m.Inner)

}
