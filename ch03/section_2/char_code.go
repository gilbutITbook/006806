package main

import "fmt"

func main() {
	var ch1 byte = 65      // 65 (10진수)
	var ch2 byte = 0101    // 65 (8진수)
	var ch3 byte = 0x41    // 65 (16진수)
	var ch4 rune = 44032   // 44032 (10진수)
	var ch5 rune = 0126000 // 44032 (8진수)
	var ch6 rune = 0xAC00  // 44032 (16진수)
	fmt.Printf("%c %c %c\n", ch1, ch2, ch3)
	fmt.Printf("%c %c %c\n", ch4, ch5, ch6)
}
