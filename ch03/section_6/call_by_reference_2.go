package main

import (
	"fmt"
	"image/color"
)

type rect struct {
	x0, y0, x1, y1 int
	color.RGBA
}

func main() {
	rect := rect{2, 4, 10, 20, color.RGBA{0xFF, 0, 0, 0xFF}}
	resize(&rect, 10, 10)
	fmt.Println(rect)
}
func resize(rect *rect, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}
