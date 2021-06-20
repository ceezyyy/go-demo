package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	p(strings.Contains("test", "es"))
	p(strings.Count("ttt", "tt"))
	p(strings.HasPrefix("ttt", "a"))
	p(strings.HasSuffix("ttt", "a"))
	p(strings.Index("ttt", "t"))
	p(strings.Join([]string{"a", "b", "c"}, "c"))
	p(strings.Repeat("a", 10))
	res := strings.Split("-a-b-c-d-e-", "-")
	p(res)
	p(res[0])
	p(strings.Replace("foo", "o", "0", -1))
	p(strings.Replace("foo", "o", "0", 1))
	p(strings.ToLower("teSt"))
	p(strings.ToUpper("ReSt"))
	p(len("hello"))
	p("hello"[0])
	// byte level
	p(len("你好"))

}
