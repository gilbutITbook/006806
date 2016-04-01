package main

import "fmt"

func main() {
	var ch int = '\u0041'
	var ch2 int = '\uAC00'
	var ch3 int = '\U00101234'
	fmt.Printf("%8d - %8d - %8d\n", ch, ch2, ch3) // 정수
	fmt.Printf("%8c - %8c - %8c\n", ch, ch2, ch3) // 문자
	fmt.Printf("%8X - %8X - %8X\n", ch, ch2, ch3) // UTF-8 바이트 수
	fmt.Printf("%8U - %8U - %8U", ch, ch2, ch3)   // UTF-8 코드값
}
