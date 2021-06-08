package main

import "fmt"

var m1 = make(map[string]int)

func main() {

	// init
	m1["monday"] = 1
	m1["tuesday"] = 2
	m1["wednesday"] = 3
	m1["thursday"] = 4
	m1["friday"] = 5
	m1["saturday"] = 6
	m1["sunday"] = 7

	// print out
	for key := range m1 {
		fmt.Printf("%s: %d\n", key, m1[key])
	}

	// test for the presence of "tuesday" and "hollyday"
	testKeyExists("tuesday")
	testKeyExists("hollyday")

}

func testKeyExists(key string) {
	if val, isPresent := m1[key]; isPresent {
		fmt.Printf("%s exists, its val is: %d\n", key, val)
	} else {
		fmt.Printf("%s does not exist\n", key)
	}
}
