package main

import (
	"fmt"
	"math"
)

func main() {

	val, err := IntFromInt64(1000000000000)
	if err != nil {
		// 直接输出错误信息
		fmt.Printf("%v", err.Error())
		return

	}
	fmt.Println(val)

}

//ConvertInt64ToInt int64 -> int
func ConvertInt64ToInt(i int64) int {
	if i <= math.MinInt32 || i >= math.MaxInt32 {
		// 异常处理
		errMsg := fmt.Sprintf("%d is out of range int32", i)
		// 抛出异常
		panic(fmt.Sprintf("Error: %s in main.ConvertInt64ToInt() with para %d", errMsg, i))
	}
	return int(i)
}

//IntFromInt64 recovers and return
// named return variable
func IntFromInt64(i int64) (val int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return ConvertInt64ToInt(i), nil
}
