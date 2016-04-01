package main

import "fmt"

type dimension struct {
	width, height, length float64
}
type Item struct {
	name             string
	price            float64
	quantity         int
	packageDimension dimension
	itemDimension    dimension
}

func main() {
	shoes := Item{
		"Sports Shoes", 30000, 2,
		dimension{30, 270, 20},
		dimension{50, 300, 30},
	}
	// 내부 필드인 dimension 구조체의 값 출력
	fmt.Printf("%#v\n", shoes.itemDimension)
	fmt.Printf("%#v\n", shoes.packageDimension)

	// dimension 구조체의 내부 필드 값 출력
	fmt.Println(shoes.packageDimension.width)
	fmt.Println(shoes.packageDimension.height)
	fmt.Println(shoes.packageDimension.length)
}
