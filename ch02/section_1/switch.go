package main

import "fmt"

func main() {
	c := 'a'
	switch {
	// case에 조건식 사용
	case '0' <= c && c <= '9':
		fmt.Printf("%c은(는) 숫자입니다", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c은(는) 소문자입니다", c)
	case 'A' <= c && c <= 'Z':
		fmt.Printf("%c은(는) 대문자입니다", c)
	}
}
