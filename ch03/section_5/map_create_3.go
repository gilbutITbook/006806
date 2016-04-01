package main

import "fmt"

func main() {
	numberMap := make(map[string]int, 3) // 용량 생략 가능
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println(numberMap)
}
