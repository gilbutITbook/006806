package main

import (
	"fmt"
	"reflect"
	"strings"
)

func TitleCase(s string) string {
	return strings.Title(s)
}
func main() {
	caption := "go is an open source programming language"
	// TitleCase를 바로 호출
	title := TitleCase(caption)
	fmt.Println(title)
	// TitleCase를 동적으로 호출
	titleFuncValue := reflect.ValueOf(TitleCase)
	values := titleFuncValue.Call([]reflect.Value{reflect.ValueOf(caption)})
	title = values[0].String()
	fmt.Println(title)
}
