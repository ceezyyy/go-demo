package main

import "fmt"

// gopl: ch2
// F = 32 + 1.8C

type C float64 // 摄氏度
type F float64 // 华氏度
type K float64 // 开尔文

const (
	AbsoluteZeroC C = -273.15
	FreezingC     C = 0
	BoilingC      C = 100
)

func main() {

	f := CtoF(BoilingC)
	fmt.Println(f)

	c := FtoC(0)
	fmt.Println(c)

	c = KtoC(0)
	fmt.Println(c)

}

func CtoF(c C) F {
	// it makes the change of meaning explicit
	return F(32 + 1.8*c)
}

func FtoC(f F) C {
	return C((f - 32) / 1.8)
}

func KtoC(k K) C {
	return C(k - 273.15)
}

func (f F) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (c C) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (k K) String() string {
	return fmt.Sprintf("%gK", k)
}
