package main

import "fmt"

type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area :", s.area())
}

type rect struct{ width, height float64 }

func (r rect) area() float64 {
	return r.width * r.height
}
func main() {
	r := rect{3, 4}
	describe(r) // area : 12
}
