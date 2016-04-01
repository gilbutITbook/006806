package main

import "fmt"

func main() {
	s1 := "hello"
	fmt.Println([]rune(s1))                              // [104 101 108 108 111]
	fmt.Println([]byte(s1))                              // [104 101 108 108 111]
	fmt.Println(string([]byte{104, 101, 108, 108, 111})) // hello
	fmt.Println(string([]rune{104, 101, 108, 108, 111})) // hello
	s2 := "안녕하세요"
	fmt.Println([]rune(s2)) // [50504 45397 54616 49464 50836]
	fmt.Println([]byte(s2)) // [236 149 136 235 133 149 237 149 152 236 132 184 236 154 148]
	// 안녕하세요
	fmt.Println(string([]rune{50504, 45397, 54616, 49464, 50836}))
	// 안녕하세요
	fmt.Println(string([]byte{236, 149, 136, 235, 133, 149, 237, 149, 152, 236, 132, 184,
		236, 154, 148}))
}
