package main

import "sync"

type InMem struct {
	sync.Mutex
	m map[string]int
}

func (i *InMem) New() *InMem {
	return &InMem{m: make(map[string]int)}
}

func (i *InMem) Get(k string) (int, bool) {
	i.Lock()
	v, ok := i.m[k]
	i.Unlock()
	return v, ok
}

func main() {
	im := InMem{}
	inMem := im.New()

	// WARNING: It should not promote data or behavior we want to hide from outside
	inMem.Lock()
}
