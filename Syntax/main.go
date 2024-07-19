package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 13

	fmt.Println("map:", m)

	v1 := m["key1"]
	fmt.Println("v1: ", v1)

	v3 := m["key3"]
	fmt.Println("v3: ", v3)

	delete(m, "key2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["key2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
