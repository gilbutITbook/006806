package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	// fmt 패키지의 Print 함수 사용
	fmt.Print("이름을 입력하세요: ")

	// fmt 패키지의 Scanf 함수 사용
	fmt.Scanf("%s", &name)

	// fmt 패키지의 Fprintf 함수 사용
	// os 패키지의 Stdout 변수 사용
	fmt.Fprintf(os.Stdout, "Hello %s\n", name)

}
