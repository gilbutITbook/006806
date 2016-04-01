package main

import "fmt"

func main() {
	i := 100000
	j := int16(10000)
	k := uint8(100)

	fmt.Println(i + j) // 컴파일 에러
	fmt.Println(i + k) // 컴파일 에러
	fmt.Println(j > k) // 컴파일 에러
}
