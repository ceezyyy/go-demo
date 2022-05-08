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
	Aggregations struct {
		Total struct {
			Value float64 `json:"value"`
		} `json:"total"`
	} `json:"aggregations"`
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

	testByte := []byte(`{"took":17,"timed_out":false,"_shards":{"total":4,"successful":4,"skipped":0,"failed":0},"hits":{"total":1,"max_score":0.0,"hits":[]},"aggregations":{"total":{"value":844.0}}}`)
	var resp Resp
	if err := json.Unmarshal(testByte, &resp); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", resp)
	fmt.Println(resp.Aggregations.Total.Value)
}
