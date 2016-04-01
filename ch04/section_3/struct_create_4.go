package main

import "fmt"

func main() {
	type rect struct{ w, h float64 }

	r1 := rect{1, 2}
	r2 := new(rect)
	r2.w, r2.h = 3, 4
	r3 := &rect{}
	r4 := &rect{5, 6}

	fmt.Println(r1) // {1 2}
	fmt.Println(r2) // &{3 4}
	fmt.Println(r3) // &{0 0}
	fmt.Println(r4) // &{5 6}
}
