package main

import "fmt"

func main() {
	numberMap := map[string]int{}
	numberMap["zero"] = 0
	numberMap["one"] = 1
	numberMap["two"] = 2

	if v, ok := numberMap["three"]; ok {
		fmt.Println("'three' is in numberMap. value: ", v)
	} else {
		fmt.Println("'three' is not in numberMap")
	}

	if _, ok := numberMap["four"]; !ok {
		numberMap["four"] = 4
	}
	fmt.Println(numberMap["four"])
}
