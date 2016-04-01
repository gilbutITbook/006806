package main

import "fmt"

func main() {
	i := 100000
	j := int16(10000)
	k := uint8(100)

	fmt.Println(i + int(j)) // 110000
	fmt.Println(i + int(k)) // 100100
}
