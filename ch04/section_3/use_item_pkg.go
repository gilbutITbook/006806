package main

import (
	"fmt"
	"gobook-example/ch04/section_3/item"
)

func main() {
	shirt := item.New("Men's Slim-Fit Shirt", 25000, 3)
	fmt.Println(shirt)

	shirt.SetPrice(30000)
	shirt.SetQuantity(2)
	fmt.Println("Name: ", shirt.Name())
	fmt.Println("Price: ", shirt.Price())
	fmt.Println("Quantity: ", shirt.Quantity())
}
