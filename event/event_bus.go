package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
)

const (
	Topic = "main:calculator"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	bus := EventBus.New()
	bus.Subscribe(Topic, calculator)
	bus.Publish(Topic, 20, 40)
}
