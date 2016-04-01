package main

import "fmt"

func main() {
	var p *int

	i := 1
	p = &i
	fmt.Println(i)  // 1
	fmt.Println(&i) // 0x1043617c
	fmt.Println(*p) // 1
	fmt.Println(p)  // 0x1043617c
}
