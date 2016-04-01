package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	var str string
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			str += s
		}
	}
	fmt.Print(str, "\n")
}

func nextString(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}
