package main

import (
	"fmt"
)

func main() {

	// 返回的是地址
	m := make(map[string]int)

	// 警惕!!!
	// But attempts to write to a nil map will cause a runtime panic!!
	//m = nil
	// 使用 len(m) == 0 来判断 map 是否为空
	fmt.Println(len(m))
	m["k1"] = 1
	m["k2"] = 2

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

	// 遍历 key
	for k := range m {
		fmt.Println(k)
	}

}
