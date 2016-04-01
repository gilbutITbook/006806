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
	var t Item
	t.name = "Men's Slim-Fit Shirt"
	t.price = 25000
	t.quantity = 3
	fmt.Println(t.name)     // Men's Slim-Fit Shirt
	fmt.Println(t.price)    // 25000
	fmt.Println(t.quantity) // 3
	fmt.Println(t.Cost())   // 75000
}
