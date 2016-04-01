package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func main() {
	p := &Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println(p)        // &{Men's Slim-Fit Shirt 25000 3}
	fmt.Println(p.Cost()) // 75000
}
