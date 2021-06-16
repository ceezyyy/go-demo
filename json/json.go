package main

import (
	"encoding/json"
	"fmt"
)

// 只有对外暴露的 field 才能被 JSON 传输
// 大写
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// basic type to JSON string
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(88)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "banana"}
	fmt.Println(slcD)
	slcB, _ := json.Marshal(slcD)
	//fmt.Printf("type: %T", slcB)
	fmt.Println(string(slcB))

	mapD := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(mapD)
	mapB, _ := json.Marshal(mapD)
	// json objects
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"a", "b", "c"}}
	fmt.Println(res1D)
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 使用 tag 自定义 json key
	res2D := &response2{
		Page:   2,
		Fruits: []string{"a", "b"},
	}
	fmt.Println(res2D)
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// json data -> go value
	byt := []byte(`{"num": 1, "strs": ["a", "b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

}
