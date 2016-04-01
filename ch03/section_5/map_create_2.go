package main

import "fmt"

func main() {
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3, // 요소를 여러 줄로 표기할 때 요소의 끝에 콤마(,)를 붙여야 함
	}
	fmt.Println(numberMap)
}
