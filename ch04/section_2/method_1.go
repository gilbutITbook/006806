package main

import "fmt"

type quantity int

func (q quantity) greaterThan(i int) bool { return int(q) > i }

func (q *quantity) increment() { *q++ }

func (q *quantity) decrement() { *q-- }

func main() {
	q := quantity(3)
	q.increment()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
}
