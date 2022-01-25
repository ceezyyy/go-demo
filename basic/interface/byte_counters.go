package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {

	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	world := "world"
	fmt.Fprintf(&c, "hello, %s", world)

	fmt.Println(c)

}
