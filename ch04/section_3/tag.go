package main

import (
	"fmt"
	"reflect"
)

type Item struct {
	name     string  "상품 이름"
	price    float64 "상품 가격"
	quantity int     "구매 수량"
}

func main() {
	tType := reflect.TypeOf(Item{})
	for i := 0; i < tType.NumField(); i++ {
		fmt.Println(tType.Field(i).Tag)
	}
}
