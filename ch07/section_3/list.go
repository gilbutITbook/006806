package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	items := list.New()
	for _, x := range strings.Split("ABCDEFGH", "") {
		items.PushBack(x)
	}
	e := items.PushFront(0)
	items.InsertAfter(1, e)
	for element := items.Front(); element != nil; element = element.Next() {
		fmt.Printf("%v ", element.Value) // 0 1 A B C D E F G H
	}
}
