package main

import (
	"encoding/json"
	"fmt"
)

// 重要!!!
// Ref: https://blog.golang.org/json

type Message struct {
	Name string
	Body string
	Time int64
}

type Resp struct {
	DmError  int64  `json:"dm_error"`
	ErrorMsg string `json:"error_msg"`
	Data     struct {
		NewDevice bool  `json:"new_device"`
		Timestamp int64 `json:"timestamp"`
	} `json:"data"`
}

// 1. json 的 key 必须是 string
// 2. pointer 指向的会被序列化
func main() {

	msg := &Message{
		Name: "hello world",
		Body: "This is body",
		Time: 123,
	}
	msgByte, _ := json.Marshal(msg)
	// 二进制, 字节数组
	fmt.Printf("%T\n", msgByte)

	var m Message
	// Unmarshal 只可以反序列化对应得上的大写字段
	_ = json.Unmarshal(msgByte, &m)
	fmt.Println(m.Name)
	fmt.Println(m.Body)
	fmt.Println(m.Time)

	testByte := []byte(`{"dm_error": 0,"error_msg": "0","data": {"new_device": true,"timestamp": 1628586405}}`)
	var resp Resp
	_ = json.Unmarshal(testByte, &resp)
	fmt.Println(resp.DmError)
	fmt.Println(resp.ErrorMsg)
	fmt.Println(resp.Data.NewDevice)
	fmt.Println(resp.Data.Timestamp)

}
