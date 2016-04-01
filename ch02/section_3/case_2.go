package main

import "fmt"

func main() {
	switch a, b := x[i], y[j]; {
	case a < b:
		fmt.Println("a는 b보다 작습니다")
	case a == b:
		fmt.Println("a와 b는 같습니다")
	case a > b:
		fmt.Println("a는 b보다 큽니다")
	}
}
