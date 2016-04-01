package main

import "fmt"

func area(w, h int) int {
	return w * h
}

func main() {
	v := area(3, 4)
	fmt.Println(v)
}
