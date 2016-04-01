package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s[0])        // 104: 문자열 s의 첫 번째 문자 'h'의 코드값
	fmt.Println(s[len(s)-1]) // 111: 문자열 s의 마지막 문자 'o'의 코드값
}
