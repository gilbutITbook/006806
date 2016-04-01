package main

import "fmt"

func enter(s string) { fmt.Println("entering:", s) }
func leave(s string) { fmt.Println("leaving:", s) }

func a() {
	enter("a")
	defer leave("a")
	fmt.Println("in a")
}

func b() {
	enter("b")
	defer leave("b")
	fmt.Println("in b")
	a()
}
func main() {
	b()
}
