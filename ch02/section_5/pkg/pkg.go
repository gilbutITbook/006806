package main

import (
	"fmt"
	"gobook-example/ch02/section_5/lib"
)

func main() {
	fmt.Println(lib.IsDigit('1')) // lib 패키지의 IsDigit 함수 사용
	fmt.Println(lib.IsDigit('a')) // lib 패키지의 IsDigit 함수 사용

	fmt.Println(lib.isSpace('\t')) // 빌드 오류 발생
}
