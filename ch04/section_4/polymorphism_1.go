package main

import "fmt"

// Coster
type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

// Item
type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

// DiscountItem
type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func main() {
	shoes := Item{"Sports Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	displayCost(shoes)      // cost:  60000
	displayCost(eventShoes) // cost:  135000
}
