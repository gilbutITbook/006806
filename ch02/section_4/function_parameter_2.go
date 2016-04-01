package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 한글 문자가 처음으로 나타나는 index 리턴
	f := func(c rune) bool {
		return unicode.Is(unicode.Hangul, c) // c가 한글문자이면 true 리턴
	}
	fmt.Println(strings.IndexFunc("Hello, 월드", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
}
