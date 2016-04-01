package main

import "fmt"

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func main() {
	r := rect{3, 4}
	fmt.Println("area :", r.area())
}
