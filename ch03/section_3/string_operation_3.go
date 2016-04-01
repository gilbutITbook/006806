package main

import (
	"bytes"
	"fmt"
	"math"
	"unicode"
)

func main() {
	var buffer bytes.Buffer
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			buffer.WriteString(s)
		}
	}
	fmt.Print(buffer.String(), "\n")
}

func nextString(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}
