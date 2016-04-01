package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s, "≡", s[:3], s[3:5], s[5:])
}
