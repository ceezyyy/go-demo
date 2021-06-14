package main

import "fmt"

func main() {

	mf := func(i obj) obj {
		// type switch
		switch v := i.(type) {
		case int:
			return v * 2
		case string:
			return v + v
		case bool:
			return !v
		default:
			return v
		}
	}

	ints := []obj{1, 2, 3}
	strings := []obj{"a", "b", "c"}
	bools := []obj{false, false}

	r1 := mapFunc(mf, ints)
	r2 := mapFunc(mf, strings)
	r3 := mapFunc(mf, bools)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)

}

// Handles value w/ unknown type
type obj interface {
}

func mapFunc(mf func(obj) obj, list []obj) (res []obj) {
	res = make([]obj, len(list))
	for i := range list {
		// apply
		res[i] = mf(list[i])
	}
	return
}
