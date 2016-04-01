package main

import "fmt"

func main() {
	sum, i := 0, 0
	for i < 10 {
		sum += i
		++i // 증감 연산자의 전치 연산은 허용하지 않음. 컴파일 에러 발생
	}
	fmt.Println(sum)
}