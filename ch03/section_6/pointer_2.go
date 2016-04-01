package main

import "fmt"

func main() {
	var p *int
	var pp **int

	i := 1
	p = &i
	pp = &p
	fmt.Println(i, *p, **pp) // 1 1 1

	i += 1
	fmt.Println(i, *p, **pp) // 2 2 2

	*p++
	fmt.Println(i, *p, **pp) // 3 3 3

	**pp++
	fmt.Println(i, *p, **pp) // 4 4 4
}
