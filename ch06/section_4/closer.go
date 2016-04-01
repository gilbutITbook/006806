package main

import (
	"fmt"
	"log"
)

type fType func(int, int) int

func errorHandler(fn fType) fType {
	return func(a int, b int) int {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)

			}
		}()
		return fn(a, b)
	}
}

func main() {
	fmt.Println(errorHandler(divide)(4, 2))
	fmt.Println(errorHandler(divide)(3, 0))
}

func divide(a int, b int) int {
	return a / b
}
