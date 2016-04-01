package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	strArr := []string{}
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			strArr = append(strArr, s)
		}
	}
	fmt.Println(strings.Join(strArr, ""))
}

func nextString(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}
