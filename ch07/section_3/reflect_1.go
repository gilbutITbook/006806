package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 1
	y := 1.1
	z := "one"
	fmt.Printf("x: %v(%v)\n", reflect.ValueOf(x).Int(), reflect.TypeOf(x))
	fmt.Printf("y: %v(%v)\n", reflect.ValueOf(y).Float(), reflect.TypeOf(y))
	fmt.Printf("z: %v(%v)\n", reflect.ValueOf(z).String(), reflect.TypeOf(z))
}
