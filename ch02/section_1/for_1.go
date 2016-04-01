package main

import "fmt"

func main() {
	sum := 0
	// for문에 초기화구문, 조건식, 후속작업 정의
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
