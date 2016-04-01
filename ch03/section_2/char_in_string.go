package main

import "fmt"

func main() {
	s1 := "hello"
	s2 := "안녕하세요"
	for i, c := range s1 {
		fmt.Printf("%c(%d)\t", c, i)
	}
	fmt.Println()
	for i, c := range s2 {
		fmt.Printf("%c(%d)\t", c, i)
	}
}
