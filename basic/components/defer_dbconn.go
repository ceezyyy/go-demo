package main

import (
	"fmt"
)

// clean-up resources
func main() {
	dbOperation()
}

func dbOperation() {
	connectDB()
	defer closeDB()
	fmt.Println("operating db")
}

func connectDB() {
	fmt.Println("connected")
}

func closeDB() {
	fmt.Println("closed")
}
