package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

type DiscountItem struct {
	Item
	discountRate float64
}

func main() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Sports Shoes", 50000, 3},
		10.00,
	}
	display(shoes.name)
	display(shoes.name, eventShoes.name)
}
func display(values ...string) {
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}
}
