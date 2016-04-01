package main

import (
	"fmt"
	"reflect"
)

func main() {
	languages := []string{"golang", "java", "c++"}
	sliceValue := reflect.ValueOf(languages)
	value := sliceValue.Index(1)
	value.SetString("ruby")
	fmt.Println(languages) // [golang ruby c++]
	x := 1
	if v := reflect.ValueOf(x); v.CanSet() {
		v.SetInt(2) // 호출되지 않음
	}
	fmt.Println(x) // 1
	v := reflect.ValueOf(&x)
	p := v.Elem()
	p.SetInt(3)
	fmt.Println(x) // 3
}
