package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 9; i++ {
		s = append(s, i)
		fmt.Printf("len: %d, cap: %d, %v\n", len(s), cap(s), s)
	}
}
