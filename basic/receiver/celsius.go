package main

import "fmt"

func main() {

	var c Celsius = 35.996
	fmt.Printf("%v\n", c)
	fmt.Printf("%f\n", c)

}

// Celsius 基于 float64 创建的新类型, 注意, 是新类型
type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("temperature: %.1f ˚C", c)
}
