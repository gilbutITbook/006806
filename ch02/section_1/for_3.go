package main

import "fmt"

func main() {
	sum, i := 0, 0
	// for문에 조건식 생략
	for {
		if i >= 10 {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}
