package main

import "fmt"

type Option struct {
	name  string
	value string
}
type Item struct {
	name     string
	price    float64
	quantity int
	Option   // 임베디드 필드
}

func main() {
	shoes := Item{"Sports Shoes", 30000, 2, Option{"color", "red"}}
	fmt.Println(shoes)
	// name 필드에 접근
	fmt.Println(shoes.name)
	// 임베디드 필드인 Option 구조체의 내부 필드 value에 접근
	fmt.Println(shoes.value)
	// 임베디드 필드인 Option 구조체의 내부 필드 name에 접근
	// Item 구조체에 이름이 같은 필드가 있으므로 Option 타입을 명시
	fmt.Println(shoes.Option.name)
}
