package main

import "fmt"

type rect struct{ width, height float64 }

func (r rect) area() float64 {
	return r.width * r.height
}

func (r *rect) resize(w, h float64) {
	r.width += w
	r.height += h
}

func main() {
	r := rect{3, 4}
	fmt.Println("area :", r.area()) // area : 12
	r.resize(10, 10)
	fmt.Println("area :", r.area()) // area : 182

	/* area() 메쏘드의 함수 표현식.
	   서명: func(rect) float64 */
	areaFn := rect.area

	/* resize() 메쏘드의 함수 표현식.
	   서명: func(*rect, float64, float64) */
	resizeFn := (*rect).resize

	fmt.Println("area :", areaFn(r)) // area : 182
	resizeFn(&r, -10, -10)
	fmt.Println("area :", areaFn(r)) // area : 12
}
