package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"time"
)

func main() {
	var (
		cacheSize = 100 * 1024 * 1024
		cache     = freecache.NewCache(cacheSize)
	)

	key := []byte("abc")
	val := []byte("def")
	expire := 60 // expire in 60 seconds
	cache.Set(key, val, expire)

	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", got)
	}

	time.Sleep(5 * time.Second)

	got, _ = cache.Get(key)
	fmt.Println(string(got))

	fmt.Println(cache.EntryCount(), cache.LookupCount(), cache.HitCount())
}
