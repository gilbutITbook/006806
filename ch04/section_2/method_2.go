package main

import "fmt"

type numberMap map[string]int

func (m numberMap) add(key string, value int) {
	m[key] = value
}

func (m numberMap) remove(key string) {
	delete(m, key)
}

func main() {
	m := make(numberMap)
	m["one"] = 1
	m["two"] = 2
	m.add("three", 3)
	fmt.Println(m) // map[one:1 two:2 three:3]
	m.remove("two")
	fmt.Println(m) // map[three:3 one:1]
}
