package main

import "fmt"

func main() {
	numberMap := make(map[string]int)
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	for k, v := range numberMap {
		fmt.Println(k, v)
	}
}
