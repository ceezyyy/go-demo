package main

import (
	"fmt"
	"net/url"
)

func main() {

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")

	prefix := "abc"
	params := v.Encode()
	urlStr := fmt.Sprintf("%s%s%s", prefix, "?", params)

	fmt.Println(urlStr)

}
