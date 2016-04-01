package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

func NewItem(name string, price float64, quantity int) *Item {
	if price <= 0 || quantity <= 0 || len(name) == 0 {
		return nil
	}
	return &Item{name, price, quantity}
}

func main() {
	t := NewItem("Men's Slim-Fit Shirt", 25000, 3)
	fmt.Println(t)
}
