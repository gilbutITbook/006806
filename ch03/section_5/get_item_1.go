package main

import "fmt"

func main() {
	numberMap := map[string]int{}
	numberMap["zero"] = 0
	numberMap["one"] = 1
	numberMap["two"] = 2
	fmt.Println(numberMap["zero"])  // 0
	fmt.Println(numberMap["one"])   // 1
	fmt.Println(numberMap["two"])   // 2
	fmt.Println(numberMap["three"]) // 0
}
