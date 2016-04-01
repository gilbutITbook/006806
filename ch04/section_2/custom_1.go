package main

import "fmt"

func main() {
	type quantity int
	type dozen []quantity

	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println(d)
}
