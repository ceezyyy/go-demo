package main

import "fmt"

//Day 基于 int 创建的新类型
type Day int

const (
	SU Day = iota
	MO
	TU
	WE
	TH
	FR
	SA
)

var arrDays = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func main() {
	fmt.Printf("%v\n", MO)
	fmt.Printf("%v\n", TU)
	fmt.Printf("%v\n", FR)
}

// 某个类型的 String 方法 -> %v 格式化输出
func (d Day) String() string {
	return fmt.Sprintf("%s", arrDays[d])
}
