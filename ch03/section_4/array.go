package main

import "fmt"

func main() {
	var a [5]int                 // 길이가 5인 int 형 배열 선언
	b := [3]int{1, 2, 3}         // 배열 선언과 동시에 값을 초기화
	c := [3]int{1, 2}            // 초깃값을 지정하지 않으면 0으로 초기화됨
	d := [...]int{4, 5, 6, 7, 8} // ...를 사용하여 배열의 길이를 지정
	e := [3][3]int{              // 다차원 배열 선언
		{1, 2, 3},
		{3, 4, 5}, // 요소를 여러 줄로 표기할 때, 요소의 마지막에 콤마를 붙여야 함
	}
	fmt.Printf("%-10T %d %v\n", a, len(a), a)
	fmt.Printf("%-10T %d %v\n", b, len(b), b)
	fmt.Printf("%-10T %d %v\n", c, len(c), c)
	fmt.Printf("%-10T %d %v\n", d, len(d), d)
	fmt.Printf("%-10T %d %v\n", e, len(e), e)
}
