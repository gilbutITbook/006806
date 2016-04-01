package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
	}

	// for문의 마지막에 세미콜론이 삽입되어 컴파일 오류가 발생
	for j := 5; j >= 0; j--
	{
		fmt.Print(j)
	}
}