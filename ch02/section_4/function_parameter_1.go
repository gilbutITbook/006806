package main

import "fmt"

func callback(y int, f func(int, int)) {
	f(y, 2) // add(1, 2)를 호출
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b) // 1 + 2 = 3
}

func main() {
	callback(1, add) // 1 + 2 = 3
}
