package main

import "fmt"

func main() {
	// 함수에 슬라이스를 전달하면 시스템의 부담을 줄일 수 있음
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	multiply(numbers, 5)
	fmt.Println(numbers)
}

func multiply(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}
