package main

import (
	"log"
	"net/http"
	"time"
)

var tracing bool

// The scope of a variable refers to the places a variable can be referenced
// Variable shadowing: a variable name declared in a block may be redeclared in an inner block
func main() {
	var c *http.Client

	if tracing {
		c, _ := createClientWithTracing()
		// the variable are used in logging calls, if not we would have compilation errors
		log.Println(c)
	} else {
		c, _ := createDefaultClient()
		log.Println(c)
	}

	// client will always be nil
	log.Println(c)
	return
}

func createClientWithTracing() (*http.Client, error) {
	c := http.DefaultClient
	c.Timeout = 1 * time.Second
	return c, nil
}

func createDefaultClient() (*http.Client, error) {
	c := http.DefaultClient
	c.Timeout = 1 * time.Second
	return c, nil
}
