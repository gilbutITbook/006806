package main

import (
	"fmt"
	"strings"
)

// Coster
type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

// Items
type Items []Coster

func (ts Items) Cost() (c float64) {
	for _, t := range ts {
		c += t.Cost()
	}
	return
}

func (ts Items) String() string {
	var s []string
	for _, t := range ts {
		s = append(s, fmt.Sprint(t))
	}
	return fmt.Sprintf("%d items. total: %.0f\n\t- %s",
		len(ts), ts.Cost(), strings.Join(s, "\n\t- "))
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

func (t Item) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

// DiscountItem
type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func (t DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", t.Item.String(), t.Cost(), t.discountRate,
		"%")
}

// Rental
type Rental struct {
	name         string
	feePerDay    float64
	periodLength int
	RentalPeriod
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays()*r.periodLength)
}

func (t Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

// Itemer
type Itemer interface {
	Coster
	fmt.Stringer
}

// Order
type Order struct {
	Itemer
	taxRate float64
}

func (o Order) Cost() float64 {
	return o.Itemer.Cost() * (1.0 + o.taxRate/100)
}
func (o Order) String() string {
	return fmt.Sprintf("Total price: %.0f(tax rate: %.2f)\n\tOrder details: %s",
		o.Cost(), o.taxRate, o.Itemer.String())
}

func main() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}

	order1 := Order{Items{shirt, eventShoes}, 10.00}
	order2 := Order{video, 5.00}

	fmt.Println(order1)
	fmt.Println(order2)
}
