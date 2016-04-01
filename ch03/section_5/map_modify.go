package main

import "fmt"

func main() {
	numberMap := map[int]string{}
	numberMap[1] = "one"
	numberMap[2] = "two"
	fmt.Println(numberMap) // map[1:one 2:two]

	numberMap[3] = "tree"
	fmt.Println(numberMap) // map[1:one 2:two 3:tree]

	numberMap[3] = "삼"
	fmt.Println(numberMap) // map[1:one 2:two 3:삼]

	delete(numberMap, 3)
	fmt.Println(numberMap) // map[1:one 2:two]
}
