package main

import "fmt"

func main() {
	type rect struct{ w, h float64 }
	var i int = 1
	var p *int = &i
	var s *rect = &rect{1, 2}
	fmt.Println(p)
	fmt.Println(s)
}
