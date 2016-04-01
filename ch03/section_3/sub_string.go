package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s[1:2]) // e
	fmt.Println(s[1:])  // ello
	fmt.Println(s[:2])  // he
	s = "안녕하세요"
	fmt.Println(s[1:2]) // ◆?
	fmt.Println(s[1:])  // ◆? ◆?녕하세요
	fmt.Println(s[:2])  // ◆
}
