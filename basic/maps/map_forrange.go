package main

import "fmt"

// map 遍历
func main() {

	// init
	capitals := make(map[string]string)
	capitals["France"] = "Paris"
	capitals["China"] = "Beijing"
	capitals["Japan"] = "Tokyo"

	// 无顺序
	for key := range capitals {
		fmt.Printf("key: %s, val: %s\n", key, capitals[key])
	}

}
