package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Learning map")

	m := make(map[string]int)
	fmt.Println("Map: ", m)

	m["key1"] = 100
	m["key2"] = 120
	m["key3"] = 108
	fmt.Println("Map:", m)

	roll := m["key3"]
	fmt.Println("Roll:", roll)

	delete(m, "key2")
	fmt.Println("Map:", m)

	value, isAvailable := m["key10"]
	fmt.Println("val:", value)
	fmt.Println("isAvailable:", isAvailable)

	clear(m)
	fmt.Println("Map:", m)

	m1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", m1)

	m2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(m1, m2) {
		fmt.Println("they are equal")
	}
}
