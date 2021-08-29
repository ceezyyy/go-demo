package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func main() {

	v := url.Values{}
	v.Set("uid", strconv.FormatInt(190903940349, 10))
	//v.Add("isTrue", "true")

	serviceURL := "serviceURL"
	// serviceURL?isTrue=true&uid=190903940349
	// serviceURL?uid=190903940349
	urlStr := fmt.Sprintf("%s%s%s", serviceURL, "?", v.Encode())

	fmt.Println(urlStr)

}
