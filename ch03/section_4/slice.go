package main

import "fmt"

func main() {
	var a []int         // int 형 슬라이스 선언. 길이와 용량은 0으로 지정됨
	b := []int{}        // int 형 슬라이스 선언. 길이와 용량은 0으로 지정됨
	c := []int{1, 2, 3} // 슬라이스 선언과 동시에 값을 초기화
	// 길이와 용량은 요소의 개수로 지정됨
	d := [][]int{ // 다차원 슬라이스 선언
		{1, 2},
		{3, 4, 5},
	}
	e := make([]int, 0)    // make 함수로 길이와 용량이 0인 슬라이스 생성
	f := make([]int, 3, 5) // make 함수로 길이가 3이고 용량이 5인 슬라이스 생성
	fmt.Printf("%-10T %d %d %v\n", a, len(a), cap(a), a)
	fmt.Printf("%-10T %d %d %v\n", b, len(b), cap(b), b)
	fmt.Printf("%-10T %d %d %v\n", c, len(c), cap(c), c)
	fmt.Printf("%-10T %d %d %v\n", d, len(d), cap(d), d)
	fmt.Printf("%-10T %d %d %v\n", e, len(e), cap(e), e)
	fmt.Printf("%-10T %d %d %v\n", f, len(f), cap(f), f)
}
