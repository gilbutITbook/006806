package main

import "fmt"

func main() {
	v := multiply(3, 4) // 컴파일 에러 발생
	fmt.Println(v)
}

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}
