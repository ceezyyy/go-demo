package main

import (
	"fmt"
	"strconv"
	"strings"
)

// strings 常用函数
func main() {

	s1 := "This is an example of string"

	fmt.Println(strings.HasPrefix(s1, "Th"))
	fmt.Println(strings.HasPrefix(s1, "Ty"))
	fmt.Println(strings.Contains(s1, "ple"))
	fmt.Println(strings.Contains(s1, "xx"))

	s2 := "Hello, world"
	fmt.Println(strings.Index(s2, "wor"))
	fmt.Println(strings.LastIndex(s2, "l"))

	s3 := "abbccc"
	// return a new string
	r1 := strings.Replace(s3, "b", "d", 2)
	fmt.Println(r1)
	// n: first n occurrences, -1 represents all
	r2 := strings.Replace(s3, "c", "d", -1)
	fmt.Println(r2)

	fmt.Println(strings.Count(s2, "l"))
	// non-overlapping
	fmt.Println(strings.Count(s3, "cc"))

	var reStr string
	reStr = strings.Repeat(s2, 2)
	fmt.Println(reStr)

	var lower string
	lower = strings.ToLower(s2)
	fmt.Println(lower)
	var upper string
	upper = strings.ToUpper(s2)
	fmt.Println(upper)

	// Trim 特指 leading & trailing
	s4 := "  testtest  "
	fmt.Println(strings.TrimSpace(s4))
	r4 := strings.Trim(s4, " t")
	fmt.Println(r4)
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))

	s5 := "I love Golang"
	// 一个或多个连续的空格
	r5 := strings.Fields(s5)
	fmt.Println(r5)

	s6 := "I-love-Golang"
	r6 := strings.Split(s6, "-")
	fmt.Println(r6)

	// atoi: ascii -> integer
	// itoa: integer -> ascii
	s7 := "666"
	var r7 int
	fmt.Println(strconv.IntSize)
	// Here, we disregard the possible error w/ blank identifier
	r7, _ = strconv.Atoi(s7)
	fmt.Println(r7)
	r7 += 1
	r8 := strconv.Itoa(r7)
	fmt.Println(r8)

}
