package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("https://gobyexample.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("status: %v\n", resp.Status)

	// read from body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
