package main

import "fmt"

func main() {

	// init
	square := &Square{side: 4}
	area := square.Area()
	fmt.Printf("%.2f", area)

}

type Shaper interface {
	// Area 计算面积
	Area() float32
}

type Square struct {
	side float32
}

// Area square 实现了 shaper 的接口
func (s *Square) Area() float32 {
	return s.side * s.side
}
