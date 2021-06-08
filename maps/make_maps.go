package main

import "fmt"

// 创建 map 方式
func main() {

	mapLit := map[string]int{"one": 1, "two": 2}
	fmt.Printf("%v\n", mapLit["one"])
	fmt.Printf("%v\n", mapLit["two"])
	fmt.Printf("%v\n", mapLit["three"])
	fmt.Printf("type: %T\n", mapLit)
	fmt.Printf("addr: %p\n", mapLit)

	mapCreated := make(map[string]float32)
	mapCreated["one"] = 1.0
	mapCreated["two"] = 2.0
	fmt.Printf("%f\n", mapCreated["one"])
	fmt.Printf("%f\n", mapCreated["two"])
	fmt.Printf("%f\n", mapCreated["three"])
	fmt.Printf("type: %T\n", mapCreated)

	mapAssigned := mapLit
	fmt.Printf("type: %T\n", mapAssigned)
	// 同一内存地址
	fmt.Printf("addr: %p\n", mapAssigned)

}
