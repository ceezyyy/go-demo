package main

import (
	"fmt"
	"git.inke.cn/BackendPlatform/go-tools/errs"
)

func main() {
	var err error
	fmt.Println(err == nil)
	if _, err := myFunc2(); err != nil {
		fmt.Println(err)
	}
}

func myFunc1() (int, error) {
	return 100, errs.New(100, "myFunc1 here")
}

func myFunc2() (res int, err error) {
	// 看看是哪个 err
	res, err = myFunc1()
	if err != nil {
		fmt.Printf("myFunc2 here!, err: %v\n", err)
		return 0, err
	}
	return
}
