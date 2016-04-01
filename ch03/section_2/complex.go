package main

import "fmt"

func main() {
	c1 := 1 + 2i            // complex128
	c2 := complex64(3 + 4i) // complex64
	c3 := complex(5, 6)     // complex128
	fmt.Println(c1, real(c1), imag(c1))
	fmt.Println(c2, real(c2), imag(c2))
	fmt.Println(c3, real(c3), imag(c3))
}
