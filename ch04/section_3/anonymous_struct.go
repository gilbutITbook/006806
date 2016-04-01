package main

import "fmt"

func main() {
	rects := []struct{ w, h int }{{1, 2}, {}, {-1, -2}}
	for _, r := range rects {
		fmt.Printf("(%d, %d) ", r.w, r.h)
	}
}
