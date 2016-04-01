package main

import "fmt"

func main() {
	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	for index, value := range numbers {
		fmt.Println(index, value)
	}
}
