package main

import "fmt"

func main() {
	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	sum := 0
	for i := range numbers {
		numbers[i] *= 2
		sum += numbers[i]
	}
	fmt.Println("numbers: ", numbers)
	fmt.Println("sun : ", sum)
}
