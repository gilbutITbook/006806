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
	item := new(Item)
	item.name = "Men's Slim-Fit Shirt"
	item.price = 25000
	item.quantity = 3

	fmt.Println(item)        // &{Men's Slim-Fit Shirt 25000 3}
	fmt.Println(item.Cost()) // 75000
}
