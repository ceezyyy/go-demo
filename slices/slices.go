package main

import "fmt"

func main() {

	//s := make([]int, 4)
	//s[0] = 0
	//s[1] = 1
	//s[2] = 2
	//fmt.Printf("addr: %p\n", s)
	//
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//}
	//
	//fmt.Println("apd")
	//s = append(s, 3)
	//for i := range s {
	//	fmt.Println(s[i])
	//}

	res := testFunc(1)
	// 返回的是地址
	fmt.Printf("%T", res)

}

type MyS struct {
	id   int
	name string
}

func testFunc(delta int) (res []*MyS) {
	return
}
